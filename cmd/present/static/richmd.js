$(() => {
    var renderer = new marked.Renderer();
    renderer.code = function (code, language) {
        if(language == 'mermaid') {
            return '<pre class="mermaid">' + code + '</pre>';
        } else if (language === "chartjs") {
            return '<div><pre class="chartjs-src">' + code + '</pre><canvas class="chartjs-canvas"/></div>'
        } else {
            return '<pre><code>'+code+'</code></pre>';
        }
    };

    $('.richmd').each(function() {
        const _this = this;
        fetch($(_this).data('url')).then(function(resp) {
            resp.text().then(function(text) {
                _this.innerHTML = marked(text, {renderer: renderer});
            })
        })
    })

    mermaid.init()
})