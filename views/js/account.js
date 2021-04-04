$(document).ready(function () {


	// $("li").hover(
	//     function() {
	//         $(this).toggleClass("active")
	// });



	// Get the modal
	var createReceiptmodal = document.getElementById("createReceiptModal");

	// Get the button that opens the modal
	var btn = document.getElementById("createReceiptBtn");

	// Get the <span> element that closes the modal
	var span = document.getElementsByClassName("close")[0];

	// When the user clicks the button, open the modal 
	btn.onclick = function () {
		createReceiptmodal.style.display = "block";
	}

	// When the user clicks on <span> (x), close the modal
	span.onclick = function () {
		createReceiptmodal.style.display = "none";
	}

	// When the user clicks anywhere outside of the modal, close it
	window.onclick = function (event) {
		if (event.target == createReceiptmodal) {
			createReceiptmodal.style.display = "none";
		}
	}




	var viewReceiptsModal = document.getElementById("viewReceiptsModal");

	// Get the button that opens the modal
	var btn = document.getElementById("viewReceiptsBtn");

	// Get the <span> element that closes the modal
	var span = document.getElementsByClassName("close")[1];

	// When the user clicks the button, open the modal 
	btn.onclick = function () {
		viewReceiptsModal.style.display = "block";
	}

	// When the user clicks on <span> (x), close the modal
	span.onclick = function () {
		viewReceiptsModal.style.display = "none";
	}

	// When the user clicks anywhere outside of the modal, close it
	window.onclick = function (event) {
		if (event.target == viewReceiptsModal) {
			viewReceiptsModal.style.display = "none";
		}
	}




	//save receipt 

	$('#saveReceiptForm')
		.ajaxForm({
			url: '/account/admin/receipt/savereceipt', // or whatever
			type: "post",
			success: function (response, textstatus, xhr) {

				if (xhr.status == 200) {




					w = window.open();
					w.document.write(response);

					//  alert(w.onload)
					w.onload = function () {
						// receiptContent.onload = function(){
						//do what you need here
						alert("loaded");
						w.print();
						w.close();
					}


				}

				createReceiptmodal.style.display = "none";
				document.location.reload()
				// document.getElementById("saveReceiptForm").reset();
				// document.getElementById('theDate').value = new Date().toISOString().substring(0, 10);

			}
		});

	///select



	// login



	var author = '<div style="position: fixed;bottom: 0;right: 20px;background-color: #fff;box-shadow: 0 4px 8px rgba(0,0,0,.05);border-radius: 3px 3px 0 0;font-size: 12px;padding: 5px 10px;"></div>';
	$("body").append(author);

	$("input[type='password'][data-eye]").each(function (i) {
		var $this = $(this),
			id = 'eye-password-' + i,
			el = $('#' + id);

		$this.wrap($("<div/>", {
			style: 'position:relative',
			id: id

		}));

		$this.css({
			paddingRight: 60
		});
		$this.after($("<div/>", {
			html: 'Show',
			class: 'btn btn-primary btn-sm',
			id: 'passeye-toggle-' + i,
		}).css({
			position: 'absolute',
			right: 10,
			top: ($this.outerHeight() / 2) - 12,
			padding: '2px 7px',
			fontSize: 12,
			cursor: 'pointer',
		}));

		$this.after($("<input/>", {
			type: 'hidden',
			id: 'passeye-' + i
		}));

		var invalid_feedback = $this.parent().parent().find('.invalid-feedback');

		if (invalid_feedback.length) {
			$this.after(invalid_feedback.clone());
		}

		$this.on("keyup paste", function () {
			$("#passeye-" + i).val($(this).val());
		});
		$("#passeye-toggle-" + i).on("click", function () {
			if ($this.hasClass("show")) {
				$this.attr('type', 'password');
				$this.removeClass("show");
				$(this).removeClass("btn-outline-primary");
			} else {
				$this.attr('type', 'text');
				$this.val($("#passeye-" + i).val());
				$this.addClass("show");
				$(this).addClass("btn-outline-primary");
			}
		});
	});

	$(".my-login-validation").submit(function () {
		var form = $(this);
		if (form[0].checkValidity() === false) {
			event.preventDefault();
			event.stopPropagation();
		}
		form.addClass('was-validated');
	});


	$("#searchReceipts").keyup(function () {
		var value = $(this).val().toLowerCase();
		$("#searchReceiptsTable tr").filter(function () {
			$(this).toggle($(this).text().toLowerCase().indexOf(value) > -1)
		});
	});




	

});


