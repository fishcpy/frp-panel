import { RootLayout } from '@/components/layout'
import { Header } from '@/components/header'
import CanvasPanel from '@/components/canvas/CanvasPanel'

export default function CanvasPage() {
  return (
      <RootLayout mainHeader={<Header />}>
        <div className="w-full">
          <CanvasPanel />
        </div>
      </RootLayout>
  )
}
