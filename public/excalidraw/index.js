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

  let SCENE_ID;
  try {
    let sc = document.getElementById("scene-id");
    if (!sc) {
      return;
    }

    SCENE_ID = sc.innerHTML.trim();
    if (!SCENE_ID || SCENE_ID == "") {
      return;
    }
  } catch (err) {
    console.log("error while reading SCENE ID");
    console.error(err);
    return;
  }

  let INITIAL_SCENE_DATA = {};

  try {
    let serverData = document
      .getElementById("initial-scene-data")
      .innerHTML.trim();

    let localData = localStorage.getItem(SCENE_ID);
    if (!localData || localData === "") {
      localData = {};
    }

    if (serverData === "") {
      INITIAL_SCENE_DATA = localData;
    } else {
      serverData = JSON.parse(serverData);
      if (
        !localData["elements"] ||
        serverData["timestamp"] > localData["timestamp"]
      ) {
        INITIAL_SCENE_DATA = serverData;
      } else {
        INITIAL_SCENE_DATA = localData;
      }
    }
  } catch (err) {
    console.error("error reading initial data");
    console.error(err);
    return;
  }

  const saveSceneData = debounce(async function (api) {
    const elems = api.getSceneElements();
    const state = api.getAppState();
    let sceneData = ExcalidrawLib.serializeAsJSON(elems, state);

    sceneData = JSON.parse(sceneData);
    sceneData["timsetamp"] = Date.now();

    const JSONSceneData = JSON.stringify(sceneData);
    try {
      localStorage.setItem(SCENE_ID, JSONSceneData);
    } catch (err) {
      console.error("failed to save to local storage");
      console.error(err);

      localStorage.clear();
    }

    const form = new FormData();
    form.append("scene", SCENE_ID);
    form.append("payload", JSONSceneData);

    navigator.sendBeacon("/update-scene-data", form);
  }, 500);

  const App = () => {
    const [excalidrawAPI, setExcalidrawAPI] = React.useState(null);
    return (
      <>
        <div style={{ height: "95vh" }}>
          <Excalidraw
            excalidrawAPI={(api) => setExcalidrawAPI(api)}
            onChange={() => {
              saveSceneData(excalidrawAPI);
            }}
            initialData={{
              elements: INITIAL_SCENE_DATA["elements"],
              appState: INITIAL_SCENE_DATA["appState"],
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
