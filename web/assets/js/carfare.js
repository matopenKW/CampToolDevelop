var direction, position;

$(function(){

    // $(document).on('click', '#btnSearch', function(){
    //     search();
    // });

    $(document).on('click', '#btnAdd', function(){
        addRow();
    });

    $(document).on('click', '.date-label', function(){
        clickRow(this);
    });

    $(document).on('change', '.data', function(){
        changeDate(this);
    });

    $(document).on('click', '.btnInsert', function(){
        insert(this);
    });

    $(document).on('click', '.btnUpdate', function(){
        update(this);
    });

    $(document).on('click', '.btnDelete', function(){
        deleteDocument(this);
    });

});

function search(){
  var $form = $('form');
  $form.attr('action', 'carfare');
  $form.attr('method', 'POST');
  $form.submit();
}

function addRow(){
    var $tr = $('#dummy tr').clone();

    var today = new Date();
    var year = today.getFullYear();
    var month = ("0"+(today.getMonth() + 1)).slice(-2);
    var date =  ("0"+today.getDate()).slice(-2)
    var targetDate = '' + year +  month + date;

    $tr.find('.date-label').html(targetDate);
    $tr.find('.date').val(targetDate);
    $('#dataTable').append($tr);

    var $addRow = $('#dataTable tr:last');
    $addRow.find('.btnInsert').show();
    $addRow.find('.btnUpdate').hide();
}

function clickRow(obj){
    var $tr = $(obj).parent('.dataTable-row');
    var open = !$tr.hasClass('select');

    var $trs = $('.dataTable-row');
    $trs.removeClass('select');
    if (open) {
        $tr.addClass('select');
    }
}

function changeDate(obj){
    var $tr = $(obj).parents('.dataTable-row');
    $tr.find('.date-label').html(obj.value);
}

function insert(obj){
    var idx = $('.btnInsert').index(obj);
    var data = getData(idx);

    var url = '/carfare:cmd/insert';
    var done = function(data){
        $('.documentId').eq(idx).val(data.documentID);
        $('.btnInsert').eq(idx).hide();
        $('.btnUpdate').eq(idx).show();
        alert('登録しました');
    }
    var fail = function(data){
        console.log(data);
        alert('登録処理に失敗しました。開き直してね★');
    }
    ajaxExecute(url, 'POST', data, done, fail);
}

function update(obj){
    var idx = $('.btnUpdate').index(obj);
    var data = getData(idx);

    var url = '/carfare:cmd/update';
    var done = function(data){
        console.log(data);
        alert('更新しました');
    }
    var fail = function(data){
        console.log(data);
        alert('更新処理に失敗しました。開き直してね★');
    }

    ajaxExecute(url, 'POST', data, done, fail);
}

function deleteDocument(obj){
    if(!confirm('削除しますがよろしくて？')){
        return;
    }

    var idx = $('.btnDelete').index(obj);
    var docId = $('.documentId').eq(idx).val();

    if (docId === ''){
        $('.dataTable-row').eq(idx).remove();

    } else {
        var url = '/carfare:cmd/delete';
        var data = {
            userId : $('#userId').val(),
            documentId : docId
        };

        var done = function(data){
            console.log(data);
            $('.dataTable-row').eq(idx).remove();
            alert('削除しました');
        }
        var fail = function(data){
            console.log(data);
            alert('削除処理に失敗しました。開き直してね★');
        }

        ajaxExecute(url, 'POST', data, done, fail);
    }
}

function getData(idx){
    var data = {};
    data['userId'] = $('#userId').val();
    data['documentId'] = $('.documentId').eq(idx).val();
    data['date'] = $('.date').eq(idx).val();
    data['start'] = $('.start').eq(idx).val();
    data['end']= $('.end').eq(idx).val();
    data['price']= $('.price').eq(idx).val();
    data['bikou']= $('.bikou').eq(idx).val();
    return data;
}
