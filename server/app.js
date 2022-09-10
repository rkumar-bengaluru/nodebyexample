var path = require('path');
var express = require('express');
var app = express();

var htmlPath = path.join(__dirname, 'html');

app.use(express.static(htmlPath));

var server = app.listen(3000, function () {
    var host = 'localhost';
    var server = app.listen(process.env.PORT || 8080, () => {
        console.log('Server is started on 127.0.0.1:'+ (process.env.PORT || 8080))
    })
});

