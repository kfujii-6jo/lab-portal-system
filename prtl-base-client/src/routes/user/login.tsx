import { createFileRoute } from '@tanstack/react-router'
import { UserLogin } from '../../pages/User/UserLogin'
export const Route = createFileRoute('/user/login')({
  component: () => <UserLogin />,
})
