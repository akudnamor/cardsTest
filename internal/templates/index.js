document.addEventListener("DOMContentLoaded", function () {
  var orders = document.querySelectorAll(".order");
  orders.forEach(function (order) {
    var min = -5;
    var max = 5;
    var randomRotation = Math.floor(Math.random() * (max - min + 1)) + min;
    order.style.transform = "rotate(" + randomRotation + "deg)";

    order.style.transition = "box-shadow 0.2s";

    order.addEventListener("mouseenter", function () {
      this.style.boxShadow = "0px 0px 40px rgb(255, 248, 5)";
    });

    order.addEventListener("mouseleave", function () {
      this.style.boxShadow = "";
    });
  });

  document.getElementById("list-btn").addEventListener("click", function () {
    window.location.href = "/list";
  });

  document.getElementById("add-btn").addEventListener("click", function () {
    window.location.href = "/add";
  });

  document.getElementById("history-btn").addEventListener("click", function () {
    window.location.href = "/history";
  });
});
