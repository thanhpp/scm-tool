import React from 'react'
import Header from '../Header/Header'

function Layout({ children }) {
    return (
        <div>
            <div className=" max-w-[1536px] mx-auto pb-[70px] bg-blue-500 text-white">
                <Header />
                <main className=''>
                    {children}
                </main>
            </div>
        </div>
    )
}

export default Layout