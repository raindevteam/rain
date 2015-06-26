var socket;

$("#Chat").on('click', function(e) {
    e.preventDefault();
    console.log('clicked');
    $.ajax({
        type: "POST",
        dataType: 'html',
        url: "/console/chat"
    }).done(function(data) {
        $('#ConsoleView ul li.selected').removeClass('selected');
        $('#ConsoleView ul li#Chat').addClass('selected');
        $('#view').html(data)
        socket = io.connect('http://localhost:8080/console')
        chat(socket);
    });
});

$("#Logs").on('click', function(e) {
    e.preventDefault();
    console.log('clicked');
    $.ajax({
        type: "POST",
        dataType: 'html',
        url: "/console/logs"
    }).done(function(data) {
        $('#ConsoleView ul li.selected').removeClass('selected');
        $('#ConsoleView ul li#Logs').addClass('selected');
        $('#view').html(data)
    });
});

$('#ConsoleView ul li.selected').click();
