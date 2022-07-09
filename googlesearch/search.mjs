var searchWeb = function() {
    var timeout = Math.floor(Math.random()*100)
    return new Promise(resolve => {
        setTimeout(function() {
            resolve("Web Result")
        },timeout);
    });
}

var searchImage = function() {
    var timeout = Math.floor(Math.random()*100)
    return new Promise(resolve => {
        setTimeout(function() {
            resolve("Image Result")
        },timeout);
    });
}

var searchVideo = function() {
    var timeout = Math.floor(Math.random()*100)
    return new Promise(resolve => {
        setTimeout(function() {
            resolve("Video Result")
        },timeout);
    });
}

export {searchWeb,searchImage,searchVideo}