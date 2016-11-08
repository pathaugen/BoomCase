

/* x */
$(document).ready(function() {
	$("#draggable").draggable();
	$('body>div').bind("dragstart", function(event, ui) {
		event.stopPropagation();
	});
});


/* x */
$(document).ready(function() {
	Feedback();
});


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
	
	
	$( ".trigger-customize-reset" ).click(function() {
		event.preventDefault();
		
		$( ".draggable" ).remove();
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
	});
});


