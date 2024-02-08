import Image from "next/image";

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between">
      
      <div className="w-full h-full bg-zinc-900 flex">
  <div className="fixed bg-zinc-900 w-[300px] h-full rounded z-10 border-r border-[#3f3f3f]">
  <h2 className="text-center mt-2 font-bold text-2xl hover:animate-pulse cursor-crosshair">CinoFilm</h2>
    <ul className="w-full my-4 text-white pl-10 mt-5">
      <div className="max-w-sm mx-auto">
            <div className="p-3 flex items-center justify-between cursor-pointer hover:bg-zinc-700">
              <div className="flex items-center">
                  <img className="rounded-full h-10 w-10" src="https://loremflickr.com/g/600/600/boy" />
                  <div className="ml-2 flex flex-col">
                      <div className="leading-snug text-sm text-slate-200 font-bold">john doe</div>
                      <div className="leading-snug text-xs text-gray-400">@john</div>
                  </div>
              </div>
          </div>
          <div className="p-3 flex items-center justify-between cursor-pointer hover:bg-zinc-700">
              <div className="flex items-center">
                  <img className="rounded-full h-10 w-10" src="https://loremflickr.com/g/600/600/boy" />
                  <div className="ml-2 flex flex-col">
                      <div className="leading-snug text-sm text-slate-200 font-bold">john doe</div>
                      <div className="leading-snug text-xs text-gray-400">@john</div>
                  </div>
              </div>
          </div>
          <div className="p-3 flex items-center justify-between cursor-pointer hover:bg-zinc-700">
              <div className="flex items-center">
                  <img className="rounded-full h-10 w-10" src="https://loremflickr.com/g/600/600/boy" />
                  <div className="ml-2 flex flex-col">
                      <div className="leading-snug text-sm text-slate-200 font-bold">john doe</div>
                      <div className="leading-snug text-xs text-gray-400">@john</div>
                  </div>
              </div>
          </div>
      </div>
    </ul>
  </div>
  <div className="absolute ml-80 w-[76%] h-full">
  <div className="h-screen flex flex-col">
    <div className="bg-black flex-1 overflow-y-scroll">
        <div className="px-4 py-2">
            <div className="flex items-center mb-2">
                <img className="w-8 h-8 rounded-full mr-2" src="https://picsum.photos/50/50" alt="User Avatar" />
                <div className="font-medium">John Doe</div>
            </div>
            <div className="bg-gray-600 rounded-lg p-2 shadow mb-2 max-w-sm">
                Hi, how can I help you?
            </div>
            <div className="flex items-center mt-2 justify-end">
                <div className="bg-blue-500 text-white rounded-lg p-2 shadow mr-2 max-w-sm">
                    Sure, I can help with that.
                </div>
                <img className="w-8 h-8 rounded-full" src="https://picsum.photos/50/50" alt="User Avatar" />
            </div>
            <div className="flex items-center mb-2">
                <img className="w-8 h-8 rounded-full mr-2" src="https://picsum.photos/50/50" alt="User Avatar" />
                <div className="font-medium">John Doe</div>
            </div>
            <div className="bg-gray-600 rounded-lg p-2 shadow mb-2 max-w-sm">
                حالت خوبه؟
            </div>
            <div className="flex items-center mb-2">
                <img className="w-8 h-8 rounded-full mr-2" src="https://picsum.photos/50/50" alt="User Avatar" />
                <div className="font-medium">John Doe</div>
            </div>
            <div className="bg-gray-600 rounded-lg p-2 shadow mb-2 max-w-sm">
                Hi, how can I help you?
            </div>
            <div className="flex items-center mb-2">
                <img className="w-8 h-8 rounded-full mr-2" src="https://picsum.photos/50/50" alt="User Avatar" />
                <div className="font-medium">John Doe</div>
            </div>
            <div className="bg-gray-600 rounded-lg p-2 shadow mb-2 max-w-sm">
                Hi, how can I help you?
            </div>
            <div className="flex items-center mt-2 justify-end">
                <div className="bg-blue-500 text-white rounded-lg p-2 shadow mr-2 max-w-sm">
                    Sure, I can help with that.
                </div>
                <img className="w-8 h-8 rounded-full" src="https://picsum.photos/50/50" alt="User Avatar" />
            </div>
            <div className="flex items-center mt-2 justify-end">
                <div className="bg-blue-500 text-white rounded-lg p-2 shadow mr-2 max-w-sm">
                    Sure, I can help with that.
                </div>
                <img className="w-8 h-8 rounded-full" src="https://picsum.photos/50/50" alt="User Avatar" />
            </div>
        </div>
    </div>
    <div className="bg-black px-4 py-2">
        <div className="flex items-center">
            <input className="w-full border rounded-full py-2 px-4 mr-2" type="text" placeholder="Type your message..." />
            <button className="bg-blue-500 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-full">
        Send
      </button>
        </div>
    </div>
</div>
  </div>

</div>
    </main>
  );
}
