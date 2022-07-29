import React, { useEffect, useState } from 'react'
import Delete from '@mui/icons-material/Delete';
import CircularProgress from '@mui/material/CircularProgress';
import Box from '@mui/material/Box';
import InboxIcon from '@mui/icons-material/Inbox';
import Pagination from '../Pagination/Pagination';
import { useDispatch, useSelector } from 'react-redux';
import ItemType from '../../components/Item/ItemType'


function ProductList() {
    // const productsData = useSelector((state) => state.products.products)
    const [itemType, setItemType] = useState()
    // const [warehouseList, setWarehouseList] = useState(warehouseData)

    const dispatch = useDispatch()


    useEffect(() => {
        const fetchProduct = async () => {
            try {
                const response = await fetch(' https://scm-tool.thanhpp.ninja/item-type')
                if (!response) {
                    throw new Error('somethign wrong');
                    return;
                }
                const data = await response.json()

                setItemType(data.data)
                // dispatch(productsSliceActions.saveProducts({
                //     data: data.data
                // }))
                // setWarehouseList(data)
                console.log(data.data)
            } catch (err) {
                console.log('fetch wrong')
            }
        }

        fetchProduct()
    }, [])

    const removeHandle = () => {

    }
    return (
        //bg-[#323259]
        <div className='bg-white rounded-md'>
            <div className='grid grid-cols-12 gap-3 text-black'>
                <div className='flex items-center col-span-1 my-[16px]'>
                    {/* <input type='checkbox' className='mr-[10px] leading-[24px] ml-[20px]' /> */}
                    <span className='mr-[10px] leading-[24px] ml-[20px]' >id</span>
                </div>
                <div className=' col-span-5 my-auto   truncate'>name</div>
                <div className=' col-span-5 my-auto truncate'>description</div>
                <div className=' col-span-1 my-auto pl-[5px] flex justify-between truncate'>
                    <span className=' mr-[10px] rounded-sm'>Delete</span>
                    <span className=' mr-[10px] rounded-sm'>Edit</span>
                </div>
            </div>
            {/* {pagination.currentData ? pagination.currentData.map((item) => <ProductItem isSelectedAll={isSelectedAll} key={item.id} id={item.id} sku={item.sku} name={item.name} category={item.category} price={item.price} amount={item.amount} vendor={item.vendor} arrivalDate={item.arrivalDate} />) : <Box sx={{ display: 'flex' }}> */}
            {itemType && itemType.map(type => <ItemType id={type.id} key={type.id} name={type.name} desc={type.desc} />)}
            {/* <CircularProgress /> */}
            {/* </Box>} */}
            {/* <Pagination getNumberPerPage={numberPerPageHandle} pageCount={pagination.pageCount} onPageChange={handleClick} /> */}
            {/* <Pagination /> */}
        </div>
    )
}

export default ProductList