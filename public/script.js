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

// TODO: Move this logic into BE
document
  .querySelector("body")
  .addEventListener("deleteDrawing", function (event) {
    const activeDrawing = document
      .getElementById("drawing-id")
      .innerHTML.trim();
    const deletedDrawing = event.detail.value;

    console.log(activeDrawing, deletedDrawing);
    if (activeDrawing === deletedDrawing) {
      let url = `/app?select-collection=${document
        .getElementById("active-collection-id")
        .innerHTML.trim()}`;
      window.location.href = url;
    }
  });

// Function to show toast
function showToast(message) {
  const toastElement = document.getElementById("toast");
  if (message && message.trim() !== "") {
    toastElement.innerText = message;
  }
  toastElement.style.display = "block";

  setTimeout(() => {
    toastElement.style.opacity = "0";
    toastElement.style.display = "none";
  }, 5000);
}

document.querySelector("body").addEventListener("showToast", (event) => {
  showToast();
});
