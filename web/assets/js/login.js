
$(function(){

	$(document).on('click', '#btnLogin', function(){
		aush();
	});

	$(document).on('click', '#forgetPassword', function(){
		alert('Comming Soon...');
	});

	$(document).on('click', '#createAccount', function(){
		showAccountRegistArea();
	});

	$(document).on('click', '#btnRegist', function(){
		createAccount();
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

function showAccountRegistArea(){
	$('#createAccount').hide();
	$('#accountParam').show();

}

function createAccount(){

	var email = $('#newMailAddress').val();
	var password = $('#newPassword').val();

	firebase.auth().createUserWithEmailAndPassword(email, password)
	.then(function(user){
		alert('ユーザーを新規登録しました。');
		login(user);
	})
	.catch(function(error) {
		var errorCode = error.code;
		var errorMessage = error.message;

		if (errorCode === "auth/email-already-in-use"){
			alert('既に登録してあるメールアドレスです。');			
		} else {
			alert(errorMessage);
		}
		console.log(error);
	});
}