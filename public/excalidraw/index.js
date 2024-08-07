function initExcalidraw() {
  const Excalidraw = window.ExcalidrawLib.Excalidraw;
  const MainMenu = window.ExcalidrawLib.MainMenu;
  const CustomMenuItem = MainMenu.ItemCustom;
  const SCENE_ID = document.getElementById("scene-id").innerHTML.trim();
  let INITIAL_SCENE_DATA = {};

  try {
    let scdata = JSON.parse(
      document.getElementById("initial-scene-data").innerHTML.trim()
    );

    INITIAL_SCENE_DATA = scdata;
  } catch (err) {
    console.error(err);
  }

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

  const saveSceneData = debounce(async function (api) {
    const elems = api.getSceneElements();
    const state = api.getAppState();
    const JSONSceneData = ExcalidrawLib.serializeAsJSON(elems, state);

    const form = new FormData();
    form.append("scene", SCENE_ID);
    form.append("payload", JSONSceneData);

    navigator.sendBeacon("/update-scene-data", form);
  }, 500);

  const App = () => {
    const [excalidrawAPI, setExcalidrawAPI] = React.useState(null);
    return (
      <>
        <div style={{ height: "80vh" }}>
          <Excalidraw
            excalidrawAPI={(api) => setExcalidrawAPI(api)}
            onChange={() => {
              saveSceneData(excalidrawAPI);
            }}
            initialData={{
              elements: INITIAL_SCENE_DATA["elements"],
              appState: INITIAL_SCENE_DATA["appState"],
            }}
          >
            <MainMenu>
              <CustomMenuItem>
                <button
                  style={{ padding: "1rem" }}
                  onClick={() => {
                    const elems = excalidrawAPI.getSceneElements();
                    const state = excalidrawAPI.getAppState();

                    console.log(elems, state);
                    const blobPromise = window.ExcalidrawLib.exportToBlob({
                      elements: elems,
                    });
                    blobPromise.then((blob) => {
                      console.log(blob);
                    });

                    console.log(ExcalidrawLib.serializeAsJSON(elems, state));
                  }}
                >
                  Save to JSON
                </button>
              </CustomMenuItem>

              <CustomMenuItem>
                <button
                  style={{ padding: "1rem" }}
                  onClick={() => {
                    let sceneData = window.prompt("Gimme JSON!");
                    sceneData = JSON.parse(sceneData);

                    excalidrawAPI.updateScene({
                      appState: sceneData["appState"],
                      elements: sceneData["elements"],
                    });
                  }}
                >
                  Load from JSON
                </button>
              </CustomMenuItem>
            </MainMenu>
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
