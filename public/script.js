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
    // defined in /excalidraw/index.js
    window.initExcalidraw();
  });

document
  .querySelector("body")
  .addEventListener("deleteScene", function (event) {
    const activeScene = document.getElementById("scene-id").innerHTML.trim();
    const deletedScene = event.detail.value;

    console.log(activeScene, deletedScene);
    if (activeScene === deletedScene) {
      let url = `/?select-collection=${document
        .getElementById("active-collection-id")
        .innerHTML.trim()}`;
      window.location.href = url;
    }
  });
