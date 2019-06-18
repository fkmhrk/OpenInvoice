$(function(){
	
	$('.actionBtn li a').tooltipster({
		theme: 'tooltipster-actionBtn'
	});
	$('.btn').tooltipster({
		theme: 'tooltipster-btn',
		//arrow: false,
		offsetY: -3
	});

	//select box customize
	$('select').easySelectBox({speed: 200});
})