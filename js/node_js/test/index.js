var express = require('express');
var util = require('util');
var app = express();

app.get('/', function(req, res) {
	res.setHeader('Content-Type', 'text/plain');
	res.end(util.format("%s - %s", new Date(), 'Got HTTP Get Request'));
});
app.listen(process.env.PORT || 80);

