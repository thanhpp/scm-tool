import Loading from '../../../components/UI/Loading'
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'
import { useSelector } from 'react-redux';


function WareHouse() {

    // const [productDetail, setProductDetail] = useState()
    const [id, setId] = useState()
    const [name, setName] = useState()
    const [desc, setDesc] = useState()
    const [location, setLocation] = useState()
    const [loading, setLoading] = useState(false)
    const router = useRouter()

    const { warehouseName } = router.query
    const [auth, setAuth] = useState()


    useEffect(() => {
        if (!localStorage.getItem('token')) {
            router.push('/login')
            return;
        }
        const token = localStorage.getItem('token')
        setAuth(token)
    }, [])

    const fetchWarehouse = async () => {
        setLoading(true)
        let token = localStorage.getItem('token')
        try {
            const res = await fetch(`https://scm-tool.thanhpp.ninja/storage`, {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            })

            if (!res.ok) {
                throw new Error('can not fetch data')
                return;
            }

            const data = await res.json()

            const item = data.data.filter((storage) => storage.id == warehouseName)[0]

            setId(item.id)
            setName(item.name)
            setDesc(item.desc)
            setLocation(item.location)

            // setProductDetail(data)

            // console.log(item.id)

            setLoading(false)
        } catch (err) {
            alert(err)
            setLoading(false)
        }
    }

    useEffect(() => {
        if (!router.isReady) {
            return;
        }

        fetchWarehouse()
    }, [router.isReady])

    const idHandle = (e) => {
        setId(e.target.value)
    }

    const nameHandle = (e) => {
        setName(e.target.value)
    }

    const descHandle = (e) => {
        setDesc(e.target.value)
    }

    const locationHandle = (e) => {
        setLocation(e.target.value)
    }

    const updateProductHandle = async () => {
        setLoading(true)

        let token = localStorage.getItem('token')

        try {
            const res = await fetch(` https://scm-tool.thanhpp.ninja/storage/${warehouseName}`, {
                method: 'PUT',
                body: JSON.stringify({
                    name,
                    desc,
                    location
                }),
                headers: {
                    'Content-type': 'application/json',
                    'Authorization': `Bearer ${token}`
                }
            })
            if (!res.ok) {
                throw new Error('something wrong');
                return;
            }

            const data = await res.json()
            if (data.error.code == 200) {
                alert(data.error.message)
                router.push('/warehouse-management')
            }
            setLoading(false)
            console.log(data)
        } catch (err) {
            alert(err)
            setLoading(false)
        }
    }

    return (!auth ? <div></div> :
        <div className='bg-blue-500 w-full h-full min-h-screen   p-[36px]'>
            <div>
                <div onClick={() => router.back()} className=' bg-white rounded-[1000px] w-[30px] h-[30px] mb-[10px] cursor-pointer'>
                    <ArrowBackIcon className='text-black pl-[4px]' />
                </div>
                <div className=' text-[28px]'>
                    {/* {name?.name} */}
                </div>
            </div>
            <div className='relative text-[#A16EFF] border-b border-white'>
                <span className='ml-[10px] cursor-pointer'>Info</span>
                <div className=' absolute w-[50px] h-[3px] bg-[#A16EFF]'></div>
            </div>
            <div className='my-[32px]'>
                <div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='id'>Id<span className='text-[#d13143]'>*</span></label>
                        </div>
                        <input id='id' type='text' value={id || ""} onChange={idHandle} className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px]' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='name'>Name<span className='text-[#d13143]'>*</span></label>
                        </div>
                        <input id='name' type='text' value={name || ""} onChange={nameHandle} className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px]' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='description'>Description<span className='text-[#d13143]'>*</span></label>
                        </div>
                        <input id='description' type='text' value={desc || ""} onChange={descHandle} className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='location'>Item type id<span className='text-[#d13143]'>*</span></label>
                        </div>
                        <input id='location' type='text' value={location || ""} onChange={locationHandle} className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div onClick={updateProductHandle} className='h-[38px] text-center hover:opacity-80 rounded-md ml-[200px] px-[12px] py-[6px] bg-transparent border border-white cursor-pointer'>
                            <button className='text-white ' >Update product</button>
                        </div>
                    </div>
                </div>
            </div>
            {loading ? <Loading /> : ''}
        </div>
    )
}

export default WareHouse