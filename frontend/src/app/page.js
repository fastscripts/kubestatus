import Image from "next/image";
import BarChart from '@/components/echarts/BarChart'
import Infobox from '@/components/Infobox'


export default function Home() {
  return (
    <>
      <main class="ml-60 pt-16 max-h-screen overflow-hidden grid-cols-2">
        <div class="px-6 py-8">
          <div class="max-w-6xl mx-auto">
            <div class="bg-white rounded-3xl p-8 mb-5">
              <h1 class="text-3xl font-bold mb-10">Welcome to KubeStatus App</h1>
              <div class="flex items-center justify-between">
                <div class="flex items-stretch">
                  <div class="text-gray-400 text-xs">Text<br />mehr Text</div>
                  <div class="h-100 border-l mx-4"></div>
                  <div class="text-gray-400 text-xs">Text<br />mehr Text</div>
                </div>
                <div class="flex items-center gap-x-2">
                  Hier könnte Text stehen
                </div>
              </div>
              <hr class="my-10" />

              <div class="grid grid-cols-3 gap-x-20">
                <div class="grid grid-cols-2 gap-4" >

                  <div class="col-span-2">
                    <div class="p-4 bg-green-100 rounded-xl drop-shadow-lg">
                      <div class="font-bold text-xl text-gray-800 leading-none">
                        Nodes im Cluster
                      </div>
                      <div class="mt-5">
                        <button type="button"
                          class="inline-flex items-center justify-center py-5 px-5 rounded-xl bg-white text-gray-800 hover:text-green-500 text-xl font-bold transition">
                          20
                        </button>
                      </div>
                    </div>
                  </div>
                  <Infobox text="das ist ein Text" number="5" color="red" />
                  <div class="p-4 bg-yellow-100 rounded-xl text-gray-800 drop-shadow-lg">
                    <div class="font-bold text-2xl leading-none">200</div>
                    <div class="mt-2">Noch irgend eine Zahl</div>
                  </div>
                  <div class="p-4 bg-yellow-100 rounded-xl text-gray-800 drop-shadow-lg">
                    <div class="font-bold text-2xl leading-none">20</div>
                    <div class="mt-2">Noch irgend eine Zahl</div>
                  </div>
                  <div class="p-4 bg-yellow-100 rounded-xl text-gray-800 drop-shadow-lg">
                    <div class="font-bold text-2xl leading-none">20</div>
                    <div class="mt-2">Noch irgend eine Zahl</div>
                  </div>
                  <div class="p-4 bg-yellow-100 rounded-xl text-gray-800 drop-shadow-lg">
                    <div class="font-bold text-2xl leading-none">20</div>
                    <div class="mt-2">Noch irgend eine Zahl</div>
                  </div>
                </div>
                <div class="col-span-2">

                </div>
              </div>

              <BarChart height="200" />
            </div>
          </div>
        </div>
      </main>

    </>

  );
}
