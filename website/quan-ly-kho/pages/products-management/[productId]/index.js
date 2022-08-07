import { useRouter } from 'next/router'
import React, { useEffect, useLayoutEffect, useState } from 'react'
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import Loading from '../../../components/UI/Loading';
import { useDispatch, useSelector } from 'react-redux';
import { productsSliceActions } from '../../../store/productsSlice';

function ProductId() {
    const [productDetail, setProductDetail] = useState()
    const [image, setImage] = useState()
    const [imageDelete, setImageDelete] = useState()
    const [loading, setLoading] = useState(false)
    const router = useRouter()

    const { productId } = router.query

    const fetchProduct = async () => {
        setLoading(true)
        try {
            const res = await fetch(`https://scm-tool.thanhpp.ninja/item`)

            if (!res.ok) {
                throw new Error('can not fetch data')
                return;
            }

            const data = await res.json()

            const item = data.data.filter((product) => product.sku === productId)

            setProductDetail(...item)
            setImage(item[0].images[0])

            const domainImage = 'https://scm-tool.thanhpp.ninja/files/'
            const domainImageLength = domainImage.length

            setImageDelete(item[0]?.images[0]?.slice(domainImageLength) || "")

            console.log(item[0]?.images[0]?.slice(domainImageLength))

            setLoading(false)
        } catch (err) {
            console.log(err)
            setLoading(false)
        }
    }

    useEffect(() => {
        if (!router.isReady) {
            return;
        }

        fetchProduct()
    }, [router.isReady])


    // const skuHandle = (e) => {
    //     // setSku(e.target.value)
    //     setProductDetail(prev => {
    //         return {
    //             ...prev,
    //             sku: e.target.value
    //         }
    //     })
    // }

    const nameHandle = (e) => {
        // setName(e.target.value)
        setProductDetail(prev => {
            return {
                ...prev,
                name: e.target.value
            }
        })
    }

    const descHandle = (e) => {
        // setDesc(e.target.value)
        setProductDetail(prev => {
            return {
                ...prev,
                desc: e.target.value
            }
        })
    }

    const itemTypeIdHandle = (e) => {
        // setItemTypeId(e.target.value)
        setProductDetail(prev => {
            return {
                ...prev,
                item_type_id: e.target.value
            }
        })
    }

    const imageHandle = (e) => {
        // setImage(e.target.value)
        console.log(e.target.files[0])
        setProductDetail(prev => {
            return {
                ...prev,
                images: e.target.files[0]
            }
        })
    }

    const updateProductHandle = async () => {
        setLoading(true)

        const dataForm = new FormData()

        // dataForm.append('sku', productDetail.sku)
        dataForm.append('name', productDetail.name)
        dataForm.append('desc', productDetail.desc)
        dataForm.append('item_type_id', productDetail.item_type_id)
        dataForm.append('new_images', productDetail.images)
        if (imageDelete) {
            dataForm.append('delete_images', imageDelete)
        }

        try {
            const res = await fetch(` https://scm-tool.thanhpp.ninja/item/${productId}`, {
                method: 'PUT',
                body: dataForm
                // headers: {
                //     "Content-type": "multipart/form-data"
                // }
            })
            if (!res.ok) {
                throw new Error('something wrong');
                return;
            }

            const data = await res.json()
            if (data.error.code == 200) {
                router.push('/products-management')
            }
            setLoading(false)
            console.log(data)
        } catch (err) {
            console.log(err)
            setLoading(false)
        }
    }

    return (
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
                            <label htmlFor='sku'>Sku<span className='text-[#d13143]'>*</span></label>
                        </div>
                        <input id='sku' type='text' onChange={() => { }} value={productDetail?.sku ?? ''} className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px]' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='name'>Name<span className='text-[#d13143]'>*</span></label>
                        </div>
                        <input id='name' type='text' value={productDetail?.name ?? ''} onChange={nameHandle} className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px]' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='description'>Description<span className='text-[#d13143]'>*</span></label>
                        </div>
                        <input id='description' type='text' value={productDetail?.desc ?? ''} onChange={descHandle} className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='itemTypeId'>Item type id<span className='text-[#d13143]'>*</span></label>
                        </div>
                        <input id='itemTypeId' type='text' value={productDetail?.item_type_id ?? ''} onChange={itemTypeIdHandle} className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' />
                    </div>
                    <div className='flex mb-[26px]'>
                        <div className='flex justify-end w-[175px] leading-[38px] mr-[20px]'>
                            <label htmlFor='image'>Image<span className='text-[#d13143]'>*</span></label>
                        </div>
                        <div>
                            <div className=' max-h-full max-w-[500px]'><img src={image} /></div>
                            <div><input id='image' type='file' onChange={imageHandle} className='text-black h-[38px] w-[380px] pl-[15px] pr-[40px] truncate' /></div>
                        </div>
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

export default ProductId