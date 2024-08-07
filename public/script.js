document.querySelector("body").addEventListener("deleteElem", function (event) {
  let node = document.getElementById(event.detail.id);
  node.remove();
});

window.dispatchEventOnBody = (detailObj) => {
  document
    .querySelector("body")
    .dispatchEvent(new CustomEvent("deleteElem", { detail: detailObj }));
};

document
  .querySelector("body")
  .addEventListener("initExcalidraw", function (event) {
    window.initExcalidraw();
  });
