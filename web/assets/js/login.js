// Your web app's Firebase configuration
var firebaseConfig = {
	apiKey: "AIzaSyCcCP5qZAP6UP_FH0iljUpmntv-EzQVbok",
	authDomain: "test2-d06da.firebaseapp.com",
	databaseURL: "https://test2-d06da.firebaseio.com",
	projectId: "test2-d06da",
	storageBucket: "test2-d06da.appspot.com",
	messagingSenderId: "399183032908",
	appId: "1:399183032908:web:d5b9d5f53c89a033a8bae9"
};

// Initialize Firebase
firebase.initializeApp(firebaseConfig);

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
	$('form').submit();
}

