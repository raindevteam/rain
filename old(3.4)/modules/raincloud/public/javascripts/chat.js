function chat(socket) {
  var $chanSelect = $('#ConsoleView #view channel select')
  var $chanSelected  = $('#ConsoleView #view channel select option:selected')
  var channel = $chanSelected.text();

  console.log(channel);

  $chanSelect.on('change', function(e) {
    channel = $('#ConsoleView #view channel select option:selected').text();
    console.log(channel);
    $('messages').fadeTo(400, 0, function() {
      // $('messages').empty();
      socket.emit('backlog', { chan: channel });
    });
  });

  socket.on('backlog', function(data) {
    console.log('got backlog')
    for(var msg in data) {
      console.log(msg)
      switch(msg) {
        case 'message':
          $('messages').append(
            '<message><name>' + data[msg.from] + '</name>' +
            '<text>' + msg.text + '</text></message><br/>'
          ); break;
        case 'action':
          $('messages').append(
            '<message><name>' + msg.from + '</name>' +
            '<action>' + msg.text + '</action></message><br/>'
          ); break;
      }
    }
    $('messages').fadeTo(400, 1);
  });

  socket.on('message', function(data) {
    console.log('got message');
    el = $('messages');
    if (data.chan == channel)
      displayMessage(data, el);
  });

  function displayMessage(data, el) {
    console.log('logging: ' + data)
    el.append(
      '<message><name>' + data.from + '</name>' +
      '<text>' + data.text + '</text></message><br/>'
    );
  }
}
