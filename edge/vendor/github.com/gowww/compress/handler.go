// Package compress provides a clever gzip compressing handler.
package compress

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"net"
	"net/http"
	"strings"
	"sync"
)

// gzippableMinSize is the minimal size (in bytes) a content needs to have to be gzipped.
//
// A TCP packet is normally 1500 bytes long.
// So if the response with the TCP headers already fits into a single packet, there will be no gain from gzip.
const gzippableMinSize = 1400

// notGzippableTypes is the list of media types having a compressed content by design.
// Gzip will not be applied to any of these content types.
//
// For best performance, only the most common officials (and future officials) are listed.
//
// Official media types: http://www.iana.org/assignments/media-types/media-types.xhtml
var notGzippableTypes = map[string]struct{}{
	"application/font-woff": {},
	"application/gzip":      {},
	"application/pdf":       {},
	"application/zip":       {},
	"audio/mp4":             {},
	"audio/mpeg":            {},
	"audio/webm":            {},
	"font/otf":              {},
	"font/ttf":              {},
	"font/woff":             {},
	"font/woff2":            {},
	"image/gif":             {},
	"image/jpeg":            {},
	"image/png":             {},
	"image/webp":            {},
	"video/h264":            {},
	"video/h265":            {},
	"video/mp4":             {},
	"video/mpeg":            {},
	"video/ogg":             {},
	"video/vp8":             {},
	"video/vp9":             {},
	"video/webm":            {},
}

var gzipPool = sync.Pool{New: func() interface{} { return gzip.NewWriter(nil) }}

// A handler provides a clever gzip compressing handler.
type handler struct {
	next http.Handler
}

// Handle returns a Handler wrapping another http.Handler.
func Handle(h http.Handler) http.Handler {
	return &handler{h}
}

// HandleFunc returns a Handler wrapping an http.HandlerFunc.
func HandleFunc(f http.HandlerFunc) http.Handler {
	return Handle(f)
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Vary", "Accept-Encoding")
	if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") || r.Header.Get("Sec-WebSocket-Key") != "" {
		h.next.ServeHTTP(w, r)
		return
	}
	cw := &compressWriter{
		ResponseWriter: w,
		gzipWriter:     gzipPool.Get().(*gzip.Writer),
		firstBytes:     new(bytes.Buffer),
	}
	defer gzipPool.Put(cw.gzipWriter)
	defer cw.close()
	h.next.ServeHTTP(cw, r)
}

// compressWriter binds the downstream response writing into gzipWriter if the first content is detected as gzippable.
type compressWriter struct {
	http.ResponseWriter
	gzipWriter  *gzip.Writer
	firstBytes  *bytes.Buffer
	status      int  // status is the buffered response status in case WriteHeader has been called downstream.
	gzipChecked bool // gzipChecked tells if the gzippable checking has been done.
	gzipUsed    bool // gzipUse tells if gzip is used for the response.
}

// WriteHeader catches a downstream WriteHeader call and buffers the status code.
// The header will be written later, at the first Write call, after the gzipping checking has been done.
func (cw *compressWriter) WriteHeader(status int) {
	cw.status = status
}

// writeBufferedHeader writes the response header when a buffered status code exists.
func (cw *compressWriter) writeBufferedHeader() {
	if cw.status > 0 {
		cw.ResponseWriter.WriteHeader(cw.status)
	}
}

// Write sets the compressing headers and calls the gzip writer, but only if the Content-Type header defines a compressible content.
// Otherwise, it calls the original Write method.
func (cw *compressWriter) Write(b []byte) (int, error) {
	if cw.gzipChecked {
		if cw.gzipUsed {
			return cw.gzipWriter.Write(b)
		}
		return cw.ResponseWriter.Write(b)
	}

	if cw.ResponseWriter.Header().Get("Content-Encoding") != "" { // Content is already encoded.
		cw.gzipCheckingDone()
		return cw.ResponseWriter.Write(b)
	}

	if cw.firstBytes.Len()+len(b) < gzippableMinSize { // Still insufficient content length to determine gzippability: buffer these first bytes.
		return cw.firstBytes.Write(b)
	}

	ct := cw.ResponseWriter.Header().Get("Content-Type")
	if ct == "" {
		ct = http.DetectContentType(append(cw.firstBytes.Bytes(), b...))
		cw.ResponseWriter.Header().Set("Content-Type", ct)
	}
	if i := strings.IndexByte(ct, ';'); i >= 0 {
		ct = ct[:i]
	}
	ct = strings.ToLower(ct)
	if _, ok := notGzippableTypes[ct]; ok {
		cw.gzipCheckingDone()
		return cw.ResponseWriter.Write(b)
	}

	cw.ResponseWriter.Header().Del("Content-Length") // Because the compressed content will have a new length.
	cw.ResponseWriter.Header().Set("Content-Encoding", "gzip")
	cw.gzipWriter.Reset(cw.ResponseWriter)
	cw.gzipUsed = true
	cw.gzipCheckingDone()
	return cw.gzipWriter.Write(b)
}

// CloseNotify implements the http.CloseNotifier interface.
// No channel is returned if CloseNotify is not implemented by an upstream response writer.
func (cw *compressWriter) CloseNotify() <-chan bool {
	n, ok := cw.ResponseWriter.(http.CloseNotifier)
	if !ok {
		return nil
	}
	return n.CloseNotify()
}

// Flush implements the http.Flusher interface.
// Nothing is done if Flush is not implemented by an upstream response writer.
func (cw *compressWriter) Flush() {
	f, ok := cw.ResponseWriter.(http.Flusher)
	if ok {
		f.Flush()
	}
}

// Hijack implements the http.Hijacker interface.
// Error http.ErrNotSupported is returned if Hijack is not implemented by an upstream response writer.
func (cw *compressWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h, ok := cw.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, http.ErrNotSupported
	}
	return h.Hijack()
}

// Push implements the http.Pusher interface.
// http.ErrNotSupported is returned if Push is not implemented by an upstream response writer or not supported by the client.
func (cw *compressWriter) Push(target string, opts *http.PushOptions) error {
	p, ok := cw.ResponseWriter.(http.Pusher)
	if !ok {
		return http.ErrNotSupported
	}
	if opts == nil {
		opts = new(http.PushOptions)
	}
	if opts.Header == nil {
		opts.Header = make(http.Header)
	}
	if enc := opts.Header.Get("Accept-Encoding"); enc == "" {
		opts.Header.Add("Accept-Encoding", "gzip")
	}
	return p.Push(target, opts)
}

// gzipCheckingDone writes the buffered data (status and firstBytes) sets the gzipChecked flag to true.
func (cw *compressWriter) gzipCheckingDone() {
	cw.writeBufferedHeader()
	if cw.gzipUsed {
		cw.firstBytes.WriteTo(cw.gzipWriter)
	} else {
		cw.firstBytes.WriteTo(cw.ResponseWriter)
	}
	cw.gzipChecked = true
}

// close writes the buffered data (status and firstBytes) if it has not been done yet and closes the gzip writer if it has been used.
func (cw *compressWriter) close() {
	if !cw.gzipChecked {
		cw.writeBufferedHeader()
		cw.firstBytes.WriteTo(cw.ResponseWriter)
	}
	if cw.gzipUsed {
		cw.gzipWriter.Close()
	}
}
