
$(function(){

	$(document).on('click', '#btnLogin', function(){

		var done = function(data) {
			console.log(data);
		}
		var fail = function(data){
			console.log(data);
		}
		ajaxExecute('/login:chkLogin', 'POST', {}, done, fail);
	});

	$(document).on('click', '#forgetPassword', function(){
		alert('Comming Soon...');
	});

	$(document).on('click', '#createAccount', function(){
		alert('Comming Soon...');
	});

});

