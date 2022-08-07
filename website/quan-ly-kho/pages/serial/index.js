import React, { useState } from 'react'
import Login from '../../components/Auth/Login'
import Loading from '../../components/UI/Loading'
import Link from 'next/link'



function SerialPage() {


    const [serial, setSerial] = useState()
    const [data, setData] = useState()
    const [loading, setLoading] = useState(false)

    const serialHandle = (e) => {
        setSerial(e.target.value)
    }

    const getSerialInforHandle = async () => {
        setLoading(true)

        try {
            const res = await fetch(`https://scm-tool.thanhpp.ninja/serial/${serial}`)

            if (!res.ok) {
                throw new Error('can not fetch data')
                return;
            }

            const data = await res.json()

            const inforData = data.data

            setData(inforData)

            console.log(inforData)

            setLoading(false)
        } catch (err) {
            console.log(err)
            setLoading(false)
        }

        setLoading(false)
    }

    return (
        <div>
            <div className='max-w-[250px] w-full mx-auto mt-[30px] mb-[30px] '>
                <Link href='/login'><div className=' text-center  p-[20px] shadow-shadowCustom font-semibold cursor-pointer '>Back to login page</div></Link>
                <div className=' text-center  p-[20px] shadow-shadowCustom font-semibold cursor-pointer '>Quét mã sản phẩm</div>
                <div className='text-center  p-[10px] shadow-shadowCustom font-semibold '><input onChange={serialHandle} className='outline-0' type="text" /></div>
                <div onClick={getSerialInforHandle} className='text-center  p-[10px] shadow-shadowCustom font-semibold cursor-pointer'><button>Get serial information</button></div>
            </div>
            {loading && <Loading />}
            {data && <table className=' w-full border-collapse border border-black'>
                <tbody>
                    <tr>
                        <th className='border border-black'>Import ticket Id</th>
                        <td className='border border-black'>{data.import_ticket_id}</td>
                    </tr>
                    <tr>
                        <th className='border border-black' rowSpan='2'>Item</th>
                        <td className='border border-black'>{data.item.name}</td>
                    </tr>
                    <tr>
                        <td className='border border-black'>{data.item.desc}</td>
                    </tr>
                    <tr>
                        <th className='border border-black' rowSpan='3'>Nft information</th>
                        <td className='border border-black'>{data.nft_info.seri}</td>
                    </tr>
                    <tr>
                        <td className='border border-black'>{data.nft_info.tx_hash}</td>
                    </tr>
                    <tr>
                        <td className='border border-black'>{data.nft_info.ipfs_cid}</td>
                    </tr>
                    <tr>
                        <th className='border border-black'>Seri</th>
                        <td className='border border-black'>{data.seri}</td>
                    </tr>
                    <tr>
                        <th className='border border-black' rowSpan='2'>Storage</th>
                        <td className='border border-black'>{data.storage.name}</td>
                    </tr>
                    <tr>
                        <td className='border border-black'>{data.storage.location}</td>
                    </tr>
                </tbody>
            </table>}
        </div>
    )
}

export default SerialPage