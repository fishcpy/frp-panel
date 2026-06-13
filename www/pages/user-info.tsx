import { RootLayout } from '@/components/layout'
import { Header } from '@/components/header'
import { UserProfileForm } from '@/components/user/user-info'

export default function UserInfoPage() {
  return (
      <RootLayout mainHeader={<Header />}>
        <div className="w-full">
          <div className="flex flex-1 flex-col">
            <div className="flex flex-1 flex-row mb-2 gap-2"></div>
          </div>
          <UserProfileForm />
        </div>
      </RootLayout>
  )
}
