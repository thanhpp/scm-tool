import '../styles/globals.css'
import Layout from '../components/Layout/Layout'
import { useRouter } from 'next/router'
import { Provider } from 'react-redux'
import store from '../store/store'
import { useSelector } from 'react-redux';


function MyApp({ Component, pageProps }) {



  const router = useRouter()
  if (router.pathname === '/login' || router.pathname === '/signup' || router.pathname === '/serial') {
    return <Provider store={store}><Component {...pageProps} /></Provider>
  } else {
    return (
      <Provider store={store}>
        <Layout>
          <Component {...pageProps} />
        </Layout>
      </Provider>
    )
  }
}

export default MyApp
