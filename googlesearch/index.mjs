import { groupCollapsed } from "console";
import express from "express"
import os from "os"
import { searchWeb, searchImage,searchVideo } from "./search.mjs";
var app = express();
var port = 8080;

app.get("/", (req,res) => {
    res.send(`hello from server ${os.hostname}`)
})

async function call() {
    var w = searchWeb()
}

app.get("/search/:query", async (req,res) => {
    const arrayOfPromises = [
        searchWeb(),
        searchImage(),
        searchVideo()
    ]
    Promise.all(arrayOfPromises.map(p => {
        return Promise.race([p,delay()])
        })).then((value) => {
            res.send(value)
    });
})
var delay = function() {
    var timeout = Math.floor(Math.random()*100)
    return new Promise(resolve => {
        setTimeout(resolve, timeout, '');
    });
}
app.listen(port, ()=> {
      
    console.log(`server listening on port ${port}`)
})