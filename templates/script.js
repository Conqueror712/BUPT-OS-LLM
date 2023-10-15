function sendRequest() {
    var content = document.getElementById("content-input").value;
  
    var data = {
      "content": content
    };
  
    fetch("/api/your-endpoint", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(data)
    })
    .then(function(response) {
        if (response.ok) {
            return response.json();
        } else {
            throw new Error("请求失败，状态码: " + response.status);
        }
    })
      
    .then(function(data) {
        // 处理响应数据
        console.log("响应数据:", data);
        // 在页面上显示响应结果
        var resultElement = document.getElementById("result");
        resultElement.textContent = data.result;
    })
    .catch(function(error) {
        console.log("请求失败:", error);
    });
}
