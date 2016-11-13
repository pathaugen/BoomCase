


/* ********** ********** ********** ********** ********** */
/* x */
$(document).ready(function() {
	$("#draggable").draggable();
	$('body>div').bind("dragstart", function(event, ui) {
		event.stopPropagation();
	});
});
/* ********** ********** ********** ********** ********** */




/* ********** ********** ********** ********** ********** */
/* x */
$(document).ready(function() {
	Feedback();
});
/* ********** ********** ********** ********** ********** */



/* ********** ********** ********** ********** ********** */
/* Customize Button - Clicking the BoomCase customize button sets off various actions */
$(document).ready(function() {
	
	
	$( ".trigger-customize" ).click(function() {
		event.preventDefault();
		$( "#case-information" ).hide( function() {  });
		$( "#case-addspeaker" ).show( function() {  });
		$( ".trigger-customize" ).hide( function() {  });
		
		/*
		$( "html" ).css( "width", "900px" );
		$( "body" ).css( "width", "900px" );
		*/
		
		$( "#customize-container" ).css( "padding", "0px" );
		$( "#case-image" ).css( "width", "100%" );
		$( "#case-image-top" ).css( "padding", "0px" );
		
		$( "#header-container" ).hide( function() {  });
		$( "#stylesheet-picker" ).hide( function() {  });

		$( "#case-title" ).hide( function() {  });

		$( "#fb-share-button" ).hide( function() {  });
		$( ".fb-share-button" ).hide( function() {  });
		$( "#footer-container" ).hide( function() {  });
		$( "#footer-container-2" ).hide( function() {  });
		$( "#featured-in" ).hide( function() {  });
		$( "#play-loud" ).hide( function() {  });
	});
	

	$( ".trigger-customize-close" ).click(function() {
		event.preventDefault();
		
		$( "#case-information" ).show( function() {  });
		$( "#case-addspeaker" ).hide( function() {  });
		$( ".trigger-customize" ).show( function() {  });
		
		$( "#customize-container" ).css( "padding", "0px 1% 0px 1%" );
		$( "#case-image" ).css( "width", "66%" );
		$( "#case-image-top" ).css( "padding", "8% 0px" );
		
		$( "#header-container" ).show( function() {  });
		$( "#stylesheet-picker" ).show( function() {  });

		$( "#case-title" ).show( function() {  });

		$( "#fb-share-button" ).show( function() {  });
		$( ".fb-share-button" ).show( function() {  });
		$( "#footer-container" ).show( function() {  });
		$( "#footer-container-2" ).show( function() {  });
		$( "#featured-in" ).show( function() {  });
		$( "#play-loud" ).show( function() {  });
		
		/*$( "#case-image-customize" ).removeProp( "font-size" );*/
		$( "#case-image-customize" ).css( "top", "auto" );
		/*$( "#case-image-customize" ).css( "transform", "translate(-50%, 0px)" );*/
		$( "#case-image-customize" ).css( "transform", "none" );
		$( "#case-image-customize" ).css( "position", "static" );
		/*$( "#case-image-customize" ).css( "bottom", "0px" );*/
	});
	
	
	/* Reset - Remove all custom drivers and remove all line items and cost from total */
	$( ".trigger-customize-reset" ).click(function() {
		event.preventDefault();
		
		/* Remove all custom drivers */
		$( ".draggable" ).remove();
		
		/* Remove all line items */
		$("#low-add div").empty();
		$("#mid-add div").empty();
		$("#high-add div").empty();
		
		/* Reset the total cost back to base cost */
		var basePrice = $("#base-price").html();
		$("#total-price").html(basePrice);
	});
	
	
	$( ".trigger-customize-save" ).click(function() {
		event.preventDefault();

		$( "#case-information" ).show( function() {  });
		$( "#case-addspeaker" ).hide( function() {  });
		$( ".trigger-customize" ).show( function() {  });
		
		$( "#customize-container" ).css( "padding", "0px 1% 0px 1%" );
		$( "#case-image" ).css( "width", "66%" );
		$( "#case-image-top" ).css( "padding", "8% 0px" );
		
		$( "#header-container" ).show( function() {  });
		$( "#stylesheet-picker" ).show( function() {  });

		$( "#case-title" ).show( function() {  });

		$( "#fb-share-button" ).show( function() {  });
		$( ".fb-share-button" ).show( function() {  });
		$( "#footer-container" ).show( function() {  });
		$( "#footer-container-2" ).show( function() {  });
		$( "#featured-in" ).show( function() {  });
		$( "#play-loud" ).show( function() {  });
		
		$( "#case-image-customize" ).css( "top", "auto" );
		$( "#case-image-customize" ).css( "transform", "none" );
		$( "#case-image-customize" ).css( "position", "static" );
	});
	
	
});
/* ********** ********** ********** ********** ********** */



/* ********** ********** ********** ********** ********** */
/* Driver Adding - Clicking a driver adds it to the speaker case to drag around */
$(document).ready(function() {
	
	var iterationSpeaker = 1;
	
	$( ".driver-info" ).click(function() {
		
		
		event.preventDefault();
		
		var imageSource = $(this).find('img').attr('src');
		var imageSize = $(this).attr('data-size');
		var imageSizeHalf = parseInt(imageSize, 10) / parseInt('2', 10);
		var imageDetails = ''; /* 'TEST: '+imageSize; */
		
		/*
		<div class="draggable" id="draggable">
			<div class="draggable-inside"><p>111</p></div>
		</div>
		*/
		var elementSpeaker = $("<div>", {
			'id'	: 'draggable-'+iterationSpeaker,
			'class'	: 'draggable',
			/*'html'	: '<div class="draggable-inside"><p>This is a test! '+iterationSpeaker+'</p></div>'*/
			'html'	: '<div class="draggable-inside"><p>'+imageDetails+'</p></div>',
			/*
			width:				200px;
			background-size:	200px;
			border-radius:		100px;
			*/
			'style'	: 'background-image: url(\''+imageSource+'\'); width: '+imageSize+'px; background-size: '+imageSize+'px; border-radius: '+imageSizeHalf+'px;'
		});
		$("#case-image-container").append(elementSpeaker);
		
		$("#draggable-"+iterationSpeaker).draggable();
		/*
		$('body>div').bind("dragstart", function(event, ui) {
			event.stopPropagation();
		});
		*/

		iterationSpeaker++;
		
		
		/* Add speaker driver to line items for price calculations */
		var driverName = $(this).find('.name-container').html();
		var driverSize = $(this).find('.inch-container').find('.size').html();
		var driverCost = $(this).find('.price-container').find('.price').html();
		var drivertype = $(this).attr('data-type');
		/* <div>INCHES" DRTIVERNAME + $COST</div> */
		var elementLineItem = $("<div>", {
			'html'	: driverSize+'" '+driverName+' + $'+driverCost
		});
		if (drivertype == "low") {
			$("#low-add").append(elementLineItem);
		}
		if (drivertype == "mid") {
			$("#mid-add").append(elementLineItem);
		}
		if (drivertype == "high") {
			$("#high-add").append(elementLineItem);
		}
		
		
		/* Add the cost of the speaker driver to the total price */
		var currentPrice = $("#total-price").html();
		var newPrice = parseInt(currentPrice, 10) + parseInt(driverCost, 10);
		$("#total-price").html(newPrice);
		
		
	});
});
/* ********** ********** ********** ********** ********** */



/* ********** ********** ********** ********** ********** */
/* Checkbox - Checking the boxes for price options adds to total cost */
$(document).ready(function() {
	$( ".checkbox-option" ).click(function() {
		if ( $(this).is(":checked") ) {
			var currentPrice = $("#total-price").html();
			var newPrice = parseInt(currentPrice, 10) + parseInt($(this).val(), 10);
			$("#total-price").html(newPrice);
		} else {
			var currentPrice = $("#total-price").html();
			var newPrice = parseInt(currentPrice, 10) - parseInt($(this).val(), 10);
			$("#total-price").html(newPrice);
		}
	});
});
/* ********** ********** ********** ********** ********** */



/* ********** ********** ********** ********** ********** */
/* Driver Sorting - Selection of various categories of drivers and hiding the rest */
$(document).ready(function() {
	
	$( "#select-drivers-all" ).css( "font-weight", "bold" );
	
	$( ".select-drivers" ).click(function() {
			event.preventDefault();
			
			$( ".select-drivers" ).css( "font-weight", "normal" );
	});
	$( "#select-drivers-all" ).click(function() {
		$( "#select-drivers-all" ).css( "font-weight", "bold" );
		
		$( ".driver-info-low" ).show( function() {  });
		$( ".driver-info-mid" ).show( function() {  });
		$( ".driver-info-high" ).show( function() {  });
	});
	$( "#select-drivers-low" ).click(function() {
		$( "#select-drivers-low" ).css( "font-weight", "bold" );
		
		$( ".driver-info-low" ).show( function() {  });
		$( ".driver-info-mid" ).hide( function() {  });
		$( ".driver-info-high" ).hide( function() {  });
	});
	$( "#select-drivers-mid" ).click(function() {
		$( "#select-drivers-mid" ).css( "font-weight", "bold" );
		
		$( ".driver-info-low" ).hide( function() {  });
		$( ".driver-info-mid" ).show( function() {  });
		$( ".driver-info-high" ).hide( function() {  });
	});
	$( "#select-drivers-high" ).click(function() {
		$( "#select-drivers-high" ).css( "font-weight", "bold" );
		
		$( ".driver-info-low" ).hide( function() {  });
		$( ".driver-info-mid" ).hide( function() {  });
		$( ".driver-info-high" ).show( function() {  });
	});
});
/* ********** ********** ********** ********** ********** */



/* ********** ********** ********** ********** ********** */
/* Dashboard Section Flipping - Changing the dashboard sections */
$(document).ready(function() {
	
	/* $( "#select-dashboard-case" ).css( "font-weight", "bold" ); */
	
	$( ".select-dashboard" ).click(function() {
			event.preventDefault();
			
			$( ".select-dashboard" ).css( "font-weight", "normal" );
	});
	$( "#select-dashboard-case" ).click(function() {
		$( "#select-dashboard-case" ).css( "font-weight", "bold" );
		
		$( "#dashboard-section-case" ).show( function() {  });
		$( "#dashboard-section-driver" ).hide( function() {  });
	});
	$( "#select-dashboard-driver" ).click(function() {
		$( "#select-dashboard-driver" ).css( "font-weight", "bold" );

		$( "#dashboard-section-case" ).hide( function() {  });
		$( "#dashboard-section-driver" ).show( function() {  });
	});
});
/* ********** ********** ********** ********** ********** */


/* ********** ********** ********** ********** ********** */
/* Admin - Add Custom Case */
$(document).ready(function() {
	
	$( "#case-container #page-formcasedriver" ).css( "display", "none" );
	
	$( "#case-container #admin-add-case" ).click(function() {
			event.preventDefault();
			
			$( "#case-container #page-formcasedriver" ).toggle();
	});
});
/* ********** ********** ********** ********** ********** */













