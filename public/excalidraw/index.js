const Excalidraw = window.ExcalidrawLib.Excalidraw;
const MainMenu = window.ExcalidrawLib.MainMenu;
const CustomMenuItem = MainMenu.ItemCustom;

const App = () => {
  const [excalidrawAPI, setExcalidrawAPI] = React.useState(null);
  return (
    <>
      <div style={{ height: "80vh" }}>
        <Excalidraw excalidrawAPI={(api) => setExcalidrawAPI(api)}>
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
