
function ajaxExecute(url, type, data, done, fail, always){
    $.ajax({
        url: url,
        type: type,
        data: data
    })
    .done(function(data){
        console.log(data);
        //var dataObj = JSON.parse(data); 
        done(data.responseJSON);
    })
    .fail(function(data){
        console.log(data);
        //var dataObj = JSON.parse(data); 
        fail(data.responseJSON);
    })
    .always(always);
}