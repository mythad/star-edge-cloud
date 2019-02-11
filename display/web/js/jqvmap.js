$(function(e){
  'use strict'

	//world map
	if ($('#world-map-gdp').length ){

		$('#world-map-gdp').vectorMap({
			map: 'world_en',
			backgroundColor: null,
			color: '#ffffff',
			hoverOpacity: 0.7,
			selectedColor: '#eff4fc',
			enableZoom: true,
			showTooltip: true,
			values: sample_data,
			scaleColors: ['#0f75ff', '#2ddcd3'],
			normalizeFunction: 'polynomial'
		});

	}

	//us map
	if ($('#usa_map').length ){

		$('#usa_map').vectorMap({
			map: 'usa_en',
			backgroundColor: null,
			color: '#ffffff',
			hoverOpacity: 0.7,
			selectedColor: '#eff4fc',
			enableZoom: true,
			showTooltip: true,
			values: sample_data,
			scaleColors: ['#2b88ff', '#00e682'],
			normalizeFunction: 'polynomial'
		});

	}
	if ($('#german').length ){
		$('#german').vectorMap({
			map : 'germany_en',
			backgroundColor: null,
			color: '#ffffff',
			hoverOpacity: 0.7,
			selectedColor: '#eff4fc',
			enableZoom: true,
			showTooltip: true,
			values: sample_data,
			scaleColors: ['#2b88ff', '#00e682'],
			normalizeFunction: 'polynomial'
		});
	}
	if ($('#russia').length ){
		$('#russia').vectorMap({
			map : 'russia_en',
			backgroundColor: null,
			color: '#ffffff',
			hoverOpacity: 0.7,
			selectedColor: '#eff4fc',
			enableZoom: true,
			showTooltip: true,
			values: sample_data,
		});
	}

});