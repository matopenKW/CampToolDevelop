
$(function(){

	$(document).on('click', '#btnLogin', function(){
		aush();
	});

	$(document).on('click', '#forgetPassword', function(){
		alert('Comming Soon...');
	});

	$(document).on('click', '#createAccount', function(){
		alert('Comming Soon...');
	});

});

function aush(){

	var email = $('#mailAddress').val();
	var password = $('#password').val();

	firebase.auth().signInWithEmailAndPassword(email, password)
	.then(function(user) {
		login(user);
	})
	.catch(function(error) {
		var errorCode = error.code;
		var errorMessage = error.message;

		if (errorCode === 'auth/wrong-password') {
		  alert('Wrong password.');
		} else {
		  alert(errorMessage);
		}
		console.log(error);
	  });
}

function login(user) {
	$('#uid').val(user.uid);
	$('#loginForm').submit();
}

