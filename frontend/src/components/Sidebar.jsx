'use client'

import { usePathname } from 'next/navigation'
import Link from 'next/link'

export function Sidebar() {
  const pathname = usePathname()

  return (
    <div class="flex flex-col justify-between h-full">
      <div class="flex-grow">
        <div class="px-4 py-6 text-center border-b">
          <h1 class="text-xl font-bold leading-none">
            <span class="text-yellow-700">KubeStatus</span> App
          </h1>
        </div>
        <div class="p-4">
          <ul class="space-y-1">
            <li>
              <Link className="{`link ${pathname === '/' ? 'active' : ''}`} flex items-center bg-gray-200 rounded-xl font-bold text-sm text-yellow-900 py-3 px-4"
                href="/">
                Home
              </Link>
            </li>
            <li>
              <Link
                className="{`link ${pathname === '/about' ? 'active' : ''}`} flex items-center bg-gray-200 rounded-xl font-bold text-sm text-yellow-900 py-3 px-4"
                href="/about" >
                About
              </Link>
            </li>
            <li>
              <a href="javascript:void(0)"
                class="flex items-center bg-gray-200 rounded-xl font-bold text-sm text-yellow-900 py-3 px-4">Menüpunkt1
              </a>
            </li>
            <li>
              <a href="javascript:void(0)"
                class="flex bg-white hover:bg-yellow-50 rounded-xl font-bold text-sm text-gray-900 py-3 px-4">Menüpinkt2
              </a>
            </li>
            <li>
              <a href="javascript:void(0)"
                class="flex bg-white hover:bg-yellow-50 rounded-xl font-bold text-sm text-gray-900 py-3 px-4">Menüpunkt3
              </a>
            </li>
            <li>
              <a href="javascript:void(0)"
                class="flex bg-white hover:bg-yellow-50 rounded-xl font-bold text-sm text-gray-900 py-3 px-4">
                Menüpunkt4
              </a>
            </li>
          </ul>
        </div>
      </div>
      <div class="p-4">
        <span class="font-bold text-sm ml-2">Nochwas</span>
      </div>
    </div>

  )
}