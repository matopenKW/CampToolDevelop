$(function(){
    $(document).on('click', '#btnLogout', function(){
        firebase.auth().signOut().then(function() {
            alert('logout!');
            $('#logoutForm').submit();
          }).catch(function(error) {
              alert('error');
          });
    });
});