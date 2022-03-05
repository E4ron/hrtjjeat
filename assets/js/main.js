function Send(method, url, data, callback) {
    let xhr = new XMLHttpRequest();
    xhr.open(method, url);
    xhr.onload = function (event) {
        callback(JSON.parse(this.response));
    };
    xhr.setRequestHeader("Content-Type", "application/json; charset=utf-8");
    xhr.send(JSON.stringify(data));
}

setInterval(() => {
    Send("GET", "/course", null, function (res) {


        console.log(res)
        let Valuates = document.querySelector("#Valuates");
        while (Valuates.children.length > 0) {
            Valuates.children[0].remove();
        }
        for (let key in res) {
            let item = document.createElement("div");
            let spanName = document.createElement("div");
            let spanValue = document.createElement("div");
            item.className = "item";
            spanName.textContent = res[key].Name;
            spanValue.textContent = res[key].Current.toFixed(2);
            item.append(spanName, spanValue);
            Valuates.append(item);
        }
    });
}, 1000);
let addValuate = document.querySelector("#AddValuate");
addValuate.onclick = function () {


    let o = {
        Current: 0,
        Min: 0,
        Max: 0,
        Name: "",
    }
    let inputs = document.querySelectorAll("#InputValuate>input")

    for (let i = 0; i < inputs.length; i++) {
        if (inputs[i].type === "number") {
            o[inputs[i].name] = +inputs[i].value;
        } else {
            o[inputs[i].name] = inputs[i].value;
        }
    }
    Send("PUT", "/course", o);

}
