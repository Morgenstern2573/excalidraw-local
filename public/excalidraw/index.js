let root;

function initExcalidraw() {
  const Excalidraw = window.ExcalidrawLib.Excalidraw;
  const MainMenu = window.ExcalidrawLib.MainMenu;
  const CustomMenuItem = MainMenu.ItemCustom;

  if (root) {
    root.unmount();
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

  let IS_EDITING = false;
  try {
    let e = document.getElementById("is-editing");
    if (!e) {
      return;
    }

    let val = e.innerHTML.trim();
    if (!val || val == "") {
      return;
    }

    if (val === "true") {
      IS_EDITING = true;
    }
  } catch (err) {
    console.log("error while reading DRAWING ID");
    console.error(err);
    return;
  }

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

    if (serverData && serverData !== "") {
      serverData = JSON.parse(serverData);
      INITIAL_DRAWING_DATA = serverData;
    }

    // let localData = localStorage.getItem(DRAWING_ID);
    // if (!localData || localData === "") {
    //   localData = {};
    // } else {
    //   localData = JSON.parse(localData);
    // }

    // if (serverData === "") {
    //   INITIAL_DRAWING_DATA = localData;
    // } else {
    //   serverData = JSON.parse(serverData);
    //   if (
    //     !localData["elements"] ||
    //     serverData["timestamp"] > localData["timestamp"]
    //   ) {
    //     INITIAL_DRAWING_DATA = serverData;
    //   } else {
    //     INITIAL_DRAWING_DATA = localData;
    //   }
    // }
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
        if (savedData === formattedDrawingData) {
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

    htmx.ajax("POST", "/app/update-drawing-data", {
      values: form,
      swap: "none",
    });
  }, 1000);

  const updateUserPosition = debounce(async function (x, y) {
    let data = new FormData();
    data.set("drawingID", DRAWING_ID);
    data.set("xPos", x);
    data.set("yPos", y);

    try {
      fetch("/app/update-user-position", {
        method: "POST",
        body: data,
      });
    } catch (err) {
      console.log("Unable to update user position");
      console.error(err);
    }
  }, 1000);

  const App = () => {
    const [excalidrawAPI, setExcalidrawAPI] = React.useState(null);
    const excalidrawAPIRef = React.useRef(null);

    React.useEffect(() => {
      excalidrawAPIRef.current = excalidrawAPI;
    }, [excalidrawAPI]);

    React.useEffect(() => {
      let intervalId;

      const cleanup = () => {
        clearInterval(intervalId);
      };

      intervalId = setInterval(() => {
        if (!excalidrawAPIRef.current) {
          return;
        }

        try {
          fetch(`/app/get-users-at-drawing?drawingID=${DRAWING_ID}`)
            .then((response) => response.json())
            .then((resp) => {
              let users = JSON.parse(resp);

              const collaborators = new Map();
              for (let user of users) {
                collaborators.set(user.userID, {
                  username: user.username,
                  pointer: {
                    x: parseFloat(user.xPos),
                    y: parseFloat(user.yPos),
                    tool: "pointer",
                  },
                });
              }

              excalidrawAPIRef.current.updateScene({
                collaborators: collaborators,
              });
            });
        } catch (err) {
          console.log("error getting user positions");
          console.error(err);
        }
      }, 1000);

      return cleanup;
    }, []);

    return (
      <>
        <div style={{ height: "90vh" }}>
          <Excalidraw
            excalidrawAPI={(api) => setExcalidrawAPI(api)}
            onChange={() => {
              saveDrawingData(excalidrawAPI);
            }}
            onPointerUpdate={(data) => {
              updateUserPosition(data.pointer.x, data.pointer.y);
            }}
            initialData={{
              elements: INITIAL_DRAWING_DATA["elements"],
              appState: INITIAL_DRAWING_DATA["appState"],
            }}
            viewModeEnabled={!IS_EDITING}
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
  root = ReactDOM.createRoot(excalidrawWrapper);
  root.render(React.createElement(App));
}

window.initExcalidraw = initExcalidraw;

initExcalidraw();
