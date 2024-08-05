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
                  const blobPromise = window.ExcalidrawLib.exportToBlob({
                    elements: elems,
                  });
                  blobPromise.then((blob) => {
                    console.log(blob);
                  });
                }}
              >
                custom item
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
