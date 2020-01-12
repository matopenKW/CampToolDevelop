
function ajaxExecute(url, type, data, done, fail, always){
    $.ajax({
        url: url,
        type: type,
        data: data
    })
    .done(function(data){
        console.log(data);
        done(data);
    })
    .fail(function(data){
        console.log(data);
        // var errMssage = data.responseJSON.errMssage;
        // if (errMssage) {
        //     alert(errMssage);
        // }
        fail(data);
    })
    .always(always);
}