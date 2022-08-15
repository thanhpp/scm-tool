import React, { useEffect, useState } from 'react'
import Loading from '../../../../components/UI/Loading'
import Link from 'next/link'
import { useRouter } from 'next/router'



function GetDatasSerialPage() {

    const [dataRender, setDataRender] = useState()
    const [loading, setLoading] = useState(false)
    const [auth, setAuth] = useState()


    const router = useRouter()
    const { pid } = router.query

    useEffect(() => {
        if (!localStorage.getItem('token')) {
            router.push('/login')
            return;
        }
        const token = localStorage.getItem('token')
        setAuth(token)
    }, [])

    useEffect(() => {
        if (!router.isReady) {
            return;
        }

        const fetchData = async () => {
            let token = localStorage.getItem('token')
            setLoading(true)
            try {
                const res = await fetch(`https://scm-tool.thanhpp.ninja/import_ticket/${pid}/serial`, {
                    headers: {
                        "Authorization": `Bearer ${token}`,
                        'Content-type': 'application/json'
                    }
                })

                if (!res.ok) {
                    throw new Error('can not fetch data')
                    return;
                }

                const data = await res.json()

                setDataRender(data.data)
                console.log(data.data)

                setLoading(false)
            } catch (err) {
                alert(err)
                setLoading(false)
            }
        }

        fetchData()

    }, [router.isReady])

    return (
        <div>
            <table className=' w-full border-collapse border border-black'>
                {dataRender && dataRender.map((item) => (
                    <tbody key={item.sku}>
                        <tr>
                            <th className='border border-black' rowSpan={item.series.length + 1}>{item.sku}</th>
                            {/* <td className='border border-black'>{'token id: '}</td> */}
                        </tr>
                        {item && item.series.map((detail) => (
                            <tr key={Math.random()}>
                                <td className='border border-black'>{detail}</td>
                            </tr>
                        ))}
                        {/* <tr>
                        <th className='border border-black'>Seri</th>
                        <td className='border border-black'>{'seri: ' + data.seri}</td>
                    </tr>
                    <tr>
                        <th className='border border-black' rowSpan='2'>Storage</th>
                        <td className='border border-black'>{'name: ' + data.storage.name}</td>
                    </tr>
                    <tr>
                        <td className='border border-black'>{'location: ' + data.storage.location}</td>
                    </tr>
                    <tr>
                        <th className='border border-black'>Supplier</th>
                        <td className='border border-black'>{'name: ' + data.supplier.name}</td>
                    </tr> */}
                    </tbody>
                ))}
            </table>
        </div>
    )
}

export default GetDatasSerialPage



