$(document).ready(function(){
  var jsonObj = JSON.parse($('#data').text());
  socket_addr = 'ws://'+ jsonObj.context + '/websocket';
  var websocket = new WebSocket(socket_addr);

  console.log("Websocket - status: " + websocket.readyState);
  websocket.onopen = function(res) {
    console.log("CONNECTION opened..." + this.readyState);
    websocket.send(jsonObj.name + ' joined the chat.');
  }
  websocket.onmessage = function(res) {
    var regEx = /(.*)~~(.*)$/;
    var dataArray = regEx.exec(res.data.replace(/\n/g,'<br/>'));
    var div_id = "div" + getRandomIntInclusive(0,50000);
    var random_color = 'rgb(' + (Math.floor(Math.random() * 256)) + ',' + (Math.floor(Math.random() * 256)) + ',' + (Math.floor(Math.random() * 256)) + ')';
    if(dataArray != null && dataArray[1] != jsonObj.name){
      $('#chat_box').append('<div id="' + div_id + '"' + ' class="messages pull-left">');
      $('#' + div_id).append('<p style="color:'+random_color +'">' + '<strong>' + dataArray[1]+ '</strong>' + '</p>');
      $('#' + div_id).append('<p>' + dataArray[2] + '</p');
      $('#chat_box').append('</div><hr/>');
      $('#notify')[0].play();

    }else if (dataArray != null && dataArray[1] == jsonObj.name) {
      $('#chat_box').append('<div id="' + div_id + '"' + ' class="mymessages pull-right">');
      $('#' + div_id).append('<p>' + dataArray[2] + '</p');
      $('#chat_box').append('</div><hr/>');
    }else{
      $('#chat_box').append('<p style="text-align:center;font-weight:bold">' + res.data + '</p><br/>');
      $('#notify')[0].play();
    }
    $('#chat_box').animate({scrollTop: $('#chat_box').prop("scrollHeight")},'fast');
  }

  websocket.onerror = function(res) {
    console.log("Error occured sending..." + m.data);
  }
  websocket.onclose = function(res) {
    console.log("Disconnected - status " + this.readyState);
  }

  var getRandomIntInclusive = function (min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
  }
  $('#chat_prompt').val('');
  $('#send').on('click',function(){
    if($('#chat_prompt').val().trim() === ""){
      return false;
    }
    websocket.send(jsonObj.name + '~~' +$('#chat_prompt').val())
    $('#chat_prompt').val('');
  });
  $('#clear_chat').on('click',function(){
    $('#chat_box').html('');
  });
  $('#leave_chat').on('click',function(){
    $('#chat_box').append('<p style="text-align:center;font-weight:bold">you left the chat.</p><br/>');
    websocket.send(jsonObj.name + ' left the chat.');
    websocket.close();
  });

  $('#chat_prompt').keypress(function (e) {
    var key = e.which;
    if(key == 13){
      $('#send').trigger('click');
      return false;
    }
  });
});
