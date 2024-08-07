function initExcalidraw() {
  const Excalidraw = window.ExcalidrawLib.Excalidraw;
  const MainMenu = window.ExcalidrawLib.MainMenu;
  const CustomMenuItem = MainMenu.ItemCustom;
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
    let scdata = document.getElementById("initial-scene-data").innerHTML.trim();
    if (scdata !== "") {
      INITIAL_SCENE_DATA = JSON.parse(scdata);
    } else {
      console.log("no initial scene data found");
    }
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
