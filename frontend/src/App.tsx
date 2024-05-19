import { useState } from "react";
import { sendPing } from "./service/pingService";

function App() {
  const [reqText, setReqText] = useState("");
  const [resText, setResText] = useState("");

  async function onSummit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    const res = await sendPing(reqText);
    setResText(res.getPing()?.getText() || "");
  }

  return (
    <>
      <form action="" onSubmit={onSummit}>
        <input
          value={reqText}
          onChange={(e) => setReqText(e.target.value)}
          type="text"
        />
        <button>Send ping</button>
      </form>

      <div>{resText}</div>
    </>
  );
}

export default App;
