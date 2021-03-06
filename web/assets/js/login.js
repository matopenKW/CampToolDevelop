
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
	.catch(firebaseErrHandring);
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
	.catch(firebaseErrHandring);
}

function firebaseErrHandring(error){
	var errorCode = error.code;
	var errorMessage = error.message;

	if (errorCode === 'auth/invalid-email') {
		alert('メールアドレスの形式が不正です。');

	} else if (errorCode === 'auth/wrong-password') {
		alert('パスワードが間違っている又は不正な形式です。');

	} else if (errorCode === 'auth/user-not-found') {
		alert('存在しないユーザー又は削除された可能性があります。');

	} else if (errorCode === 'auth/email-already-in-use'){
		alert('既に登録してあるメールアドレスです。');

	} else if (errorCode === 'auth/weak-password') {
		alert('パスワードは６桁以上で登録してください。');

	} else {
		  alert(errorMessage);

	}

	console.log(error);

}