(function($) {
    "use strict";
	$(document).ready(function() {
		$('#sidebarCollapse').on('click', function() {
			$('#sidebar').toggleClass('active');
			$('body').toggleClass('toggle');
		});
		Waves.init();
		Waves.attach('.wave-effect', ['waves-button']);
		Waves.attach('.wave-effect-float', ['waves-button', 'waves-float']);
	});
	$(function() {
		$('.slimescroll-id').slimScroll({
			height: 'auto'
		});
	});
	 $( ".cover-image").each(function() {
		  var attr = $(this).attr('data-image-src');
		
		  if (typeof attr !== typeof undefined && attr !== false) {
			  $(this).css('background', 'url('+attr+') center center');
		  }
	});
	$(window).on("load", function(e) {
		$("#loading").fadeOut("slow");
	})
	$(document).ready(function() {
		$("#sidebar a").each(function() {
		  var pageUrl = window.location.href.split(/[?#]/)[0];
			if (this.href == pageUrl) {
				$(this).addClass("active");
				$(this).parent().addClass("active"); // add active to li of the current link
				$(this).parent().parent().prev().addClass("active"); // add active class to an anchor
				$(this).parent().parent().prev().click(); // click the item to make it drop
			}
		});
	});

    // ______________Full screen URL: http://www.bootstrapmb.com
	$("#fullscreen-button").on("click", function toggleFullScreen() {
      if ((document.fullScreenElement !== undefined && document.fullScreenElement === null) || (document.msFullscreenElement !== undefined && document.msFullscreenElement === null) || (document.mozFullScreen !== undefined && !document.mozFullScreen) || (document.webkitIsFullScreen !== undefined && !document.webkitIsFullScreen)) {
        if (document.documentElement.requestFullScreen) {
          document.documentElement.requestFullScreen();
        } else if (document.documentElement.mozRequestFullScreen) {
          document.documentElement.mozRequestFullScreen();
        } else if (document.documentElement.webkitRequestFullScreen) {
          document.documentElement.webkitRequestFullScreen(Element.ALLOW_KEYBOARD_INPUT);
        } else if (document.documentElement.msRequestFullscreen) {
          document.documentElement.msRequestFullscreen();
        }
      } else {
        if (document.cancelFullScreen) {
          document.cancelFullScreen();
        } else if (document.mozCancelFullScreen) {
          document.mozCancelFullScreen();
        } else if (document.webkitCancelFullScreen) {
          document.webkitCancelFullScreen();
        } else if (document.msExitFullscreen) {
          document.msExitFullscreen();
        }
      }
    })

	
	

	// ______________ BACK TO TOP BUTTON

	$(window).on("scroll", function(e) {
    	if ($(this).scrollTop() > 0) {
            $('#back-to-top').fadeIn('slow');
        } else {
            $('#back-to-top').fadeOut('slow');
        }
    });

    $("#back-to-top").on("click", function(e){
        $("html, body").animate({
            scrollTop: 0
        }, 600);
        return false;
    });
	var ratingOptions = {
		selectors: {
			starsSelector: '.rating-stars',
			starSelector: '.rating-star',
			starActiveClass: 'is--active',
			starHoverClass: 'is--hover',
			starNoHoverClass: 'is--no-hover',
			targetFormElementSelector: '.rating-value'
		}
	};
	$(".rating-stars").ratingStars(ratingOptions);
	$(".vscroll").mCustomScrollbar();
	$(".nav-sidebar").mCustomScrollbar({
		theme:"minimal",
		autoHideScrollbar: true,
		scrollbarPosition: "outside"
	});
	if ($('.chart-circle').length) {
		$('.chart-circle').each(function() {
			let $this = $(this);

			$this.circleProgress({
			  fill: {
				color: $this.attr('data-color')
			  },
			  size: $this.height(),
			  startAngle: -Math.PI / 4 * 2,
			  emptyFill: '#f9faff',
			  lineCap: ''
			});
		});
	  }
	  
	// searching toggle
	var sp = document.querySelector('.search-open');
	var searchbar = document.querySelector('.search-inline');
	var shclose = document.querySelector('.search-close');
	function changeClass() {
		searchbar.classList.add('search-visible');
	}
	function closesearch() {
		searchbar.classList.remove('search-visible');
	}
	sp.addEventListener('click', changeClass);
	shclose.addEventListener('click', closesearch);
	
})(jQuery);

$(function(e) {
		  /** Constant div card */
	  const DIV_CARD = 'div.card';
	  /** Initialize tooltips */
	  $('[data-toggle="tooltip"]').tooltip();

	  /** Initialize popovers */
	  $('[data-toggle="popover"]').popover({
		html: true
	  });
			 /** Function for remove card */
	  $('[data-toggle="card-remove"]').on('click', function(e) {
		let $card = $(this).closest(DIV_CARD);

		$card.remove();

		e.preventDefault();
		return false;
	  });

	  /** Function for collapse card */
	  $('[data-toggle="card-collapse"]').on('click', function(e) {
		let $card = $(this).closest(DIV_CARD);

		$card.toggleClass('card-collapsed');

		e.preventDefault();
		return false;
	  });
	  $('[data-toggle="card-fullscreen"]').on('click', function(e) {
		let $card = $(this).closest(DIV_CARD);

		$card.toggleClass('card-fullscreen').removeClass('card-collapsed');

		e.preventDefault();
		return false;
	  });
  });


