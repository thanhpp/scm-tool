import React, { useState } from 'react'
import Login from '../../components/Auth/Login'
import Loading from '../../components/UI/Loading'
import Link from 'next/link'



function LoginPage() {


    // const [serial, setSerial] = useState()
    const [loading, setLoading] = useState(false)

    // const serialHandle = (e) => {
    //     setSerial(e.target.value)
    // }

    // const getSerialInforHandle = async () => {
    //     setLoading(true)

    //     try {
    //         const res = await fetch(`https://scm-tool.thanhpp.ninja/serial/${serial}`)

    //         if (!res.ok) {
    //             throw new Error('can not fetch data')
    //             return;
    //         }

    //         const data = await res.json()

    //         const inforData = data.data

    //         console.log(inforData)




    //         console.log(data.data)

    //         setLoading(false)
    //     } catch (err) {
    //         console.log(err)
    //         setLoading(false)
    //     }

    //     setLoading(false)
    // }

    return (
        <div>
            <Login />
            <div className='max-w-[250px] w-full mx-auto mt-[30px] mb-[30px] '>
                <div className=' text-center  p-[20px] mb-3 '>Dành cho khách hàng</div>
                <Link href='/serial'><div className=' text-center  p-[20px] shadow-shadowCustom font-semibold cursor-pointer '>Quét mã sản phẩm</div></Link>
                {/* <div className='text-center  p-[10px] shadow-shadowCustom font-semibold '><input onChange={serialHandle} className='outline-0' type="text" /></div>
                <div onClick={getSerialInforHandle} className='text-center  p-[10px] shadow-shadowCustom font-semibold cursor-pointer'><button>Get serial information</button></div> */}
            </div>
            {loading && <Loading />}
        </div>
    )
}

export default LoginPage