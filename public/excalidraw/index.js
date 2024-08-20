function initExcalidraw() {
  const Excalidraw = window.ExcalidrawLib.Excalidraw;
  const MainMenu = window.ExcalidrawLib.MainMenu;
  const CustomMenuItem = MainMenu.ItemCustom;

  const debounce = (fn, timeout) => {
    let handle = 0;
    let lastArgs = null;

    const retv = function (...args) {
      lastArgs = args;
      clearTimeout(handle);

      handle = setTimeout(() => {
        lastArgs = null;
        fn(...args);
      }, timeout);
    };

    retv.flush = function () {
      clearTimeout(handle);
      if (lastArgs) {
        const _lastArgs = lastArgs;
        lastArgs = null;
        fn(..._lastArgs);
      }
    };

    retv.cancel = function () {
      lastArgs = null;
      clearTimeout(handle);
    };

    return retv;
  };

  let DRAWING_ID;
  try {
    let sc = document.getElementById("drawing-id");
    if (!sc) {
      return;
    }

    DRAWING_ID = sc.innerHTML.trim();
    if (!DRAWING_ID || DRAWING_ID == "") {
      return;
    }
  } catch (err) {
    console.log("error while reading DRAWING ID");
    console.error(err);
    return;
  }

  let INITIAL_DRAWING_DATA = {};

  try {
    let serverData = document
      .getElementById("initial-drawing-data")
      .innerHTML.trim();

    let localData = localStorage.getItem(DRAWING_ID);
    if (!localData || localData === "") {
      localData = {};
    }

    if (serverData === "") {
      INITIAL_DRAWING_DATA = localData;
    } else {
      serverData = JSON.parse(serverData);
      if (
        !localData["elements"] ||
        serverData["timestamp"] > localData["timestamp"]
      ) {
        INITIAL_DRAWING_DATA = serverData;
      } else {
        INITIAL_DRAWING_DATA = localData;
      }
    }
  } catch (err) {
    console.error("error reading initial data");
    console.error(err);
    return;
  }

  const saveDrawingData = debounce(async function (api) {
    const elems = api.getSceneElements();
    const state = api.getAppState();
    let drawingData = ExcalidrawLib.serializeAsJSON(elems, state);

    try {
      let savedData = localStorage.getItem(DRAWING_ID);
      if (savedData) {
        savedData = JSON.parse(savedData);
        delete savedData["timestamp"];

        savedData = JSON.stringify(savedData);
        let formattedDrawingData = drawingData.replace(/\s/g, "");
        console.log(savedData, formattedDrawingData);
        if (savedData === formattedDrawingData) {
          console.log("early return");
          return;
        }
      }
    } catch (err) {
      console.error(err);
    }

    drawingData = JSON.parse(drawingData);
    drawingData["timestamp"] = Date.now();

    const JSONDrawingData = JSON.stringify(drawingData);
    try {
      localStorage.setItem(DRAWING_ID, JSONDrawingData);
    } catch (err) {
      console.error("failed to save to local storage");
      console.error(err);

      localStorage.clear();
    }

    const form = new FormData();
    form.append("drawing", DRAWING_ID);
    form.append("payload", JSONDrawingData);

    navigator.sendBeacon("/app/update-drawing-data", form);
  }, 1000);

  const App = () => {
    const [excalidrawAPI, setExcalidrawAPI] = React.useState(null);
    return (
      <>
        <div style={{ height: "95vh" }}>
          <Excalidraw
            excalidrawAPI={(api) => setExcalidrawAPI(api)}
            onChange={() => {
              saveDrawingData(excalidrawAPI);
            }}
            initialData={{
              elements: INITIAL_DRAWING_DATA["elements"],
              appState: INITIAL_DRAWING_DATA["appState"],
            }}
            autoFocus={true}
          >
            {/* TODO: Remove excalidraw socials from main menu, leave the rest */}
            {/* <MainMenu>
            </MainMenu> */}
          </Excalidraw>
        </div>
      </>
    );
  };

  const excalidrawWrapper = document.getElementById("app");
  const root = ReactDOM.createRoot(excalidrawWrapper);
  root.render(React.createElement(App));
}

window.initExcalidraw = initExcalidraw;

initExcalidraw();
