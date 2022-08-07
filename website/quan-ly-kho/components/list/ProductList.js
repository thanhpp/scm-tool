import React, { useEffect, useState } from 'react'
import Delete from '@mui/icons-material/Delete';
import CircularProgress from '@mui/material/CircularProgress';
import Box from '@mui/material/Box';
import InboxIcon from '@mui/icons-material/Inbox';
import Pagination from '../Pagination/Pagination';
import { useSelector } from 'react-redux';
import { productsSliceActions } from '../../store/productsSlice';
import ItemProduct from '../Item/ItemProduct';



function ProductList() {
    const [pagination, setPagination] = useState({
        offset: 0,
        numberPerPage: 2,
        pageCount: 0,
    })

    const [productList, setProductList] = useState()

    const productSelector = useSelector((state) => state.products.products)
    const isSearchingSelector = useSelector((state) => state.products.isSearching)


    useEffect(() => {
        const fetchProduct = async () => {
            try {
                const response = await fetch(' https://scm-tool.thanhpp.ninja/item')
                if (!response) {
                    throw new Error('somethign wrong');
                    return;
                }
                const data = await response.json()

                setPagination(prev => ({
                    ...prev,
                    pageCount: !isSearchingSelector ? Math.ceil(data.data.length / prev.numberPerPage) : Math.ceil(productSelector.length / prev.numberPerPage)
                }))

                const productDisplay = !isSearchingSelector ? data.data.slice(pagination.offset, parseInt(pagination.offset) + parseInt(pagination.numberPerPage)) : productSelector.slice(pagination.offset, parseInt(pagination.offset) + parseInt(pagination.numberPerPage))

                setProductList(productDisplay)
                console.log(productDisplay)
            } catch (err) {
                console.log('fetch wrong')
            }
        }

        fetchProduct()
    }, [pagination.offset, pagination.numberPerPage, pagination.pageCount, isSearchingSelector, productSelector])

    const numberPerPageHandle = (value) => {
        setPagination(prev => ({
            ...prev,
            numberPerPage: value
        }))
    }

    const handleClick = (e) => {
        const selected = e.selected
        const offset = selected * pagination.numberPerPage
        setPagination({ ...pagination, offset })
    }


    return (
        //bg-[#323259]
        <div className='bg-white rounded-md'>
            <div className='grid grid-cols-12 gap-3 text-black'>
                <div className='flex items-center col-span-1 my-[16px]'>
                    {/* <input type='checkbox' className='mr-[10px] leading-[24px] ml-[20px]' /> */}
                    <span className='mr-[10px] leading-[24px] ml-[20px]' >sku</span>
                </div>
                <div className=' col-span-3 my-auto   truncate'>name</div>
                <div className=' col-span-5 my-auto truncate'>description</div>
                <div className=' col-span-2 my-auto truncate'>item type id</div>
                <div className=' col-span-1 my-auto pl-[5px] flex justify-between truncate'>
                    <span className=' mr-[10px] rounded-sm'>Delete</span>
                    <span className=' mr-[10px] rounded-sm'>Edit</span>
                </div>
            </div>
            {/* {pagination.currentData ? pagination.currentData.map((item) => <ProductItem isSelectedAll={isSelectedAll} key={item.id} id={item.id} sku={item.sku} name={item.name} category={item.category} price={item.price} amount={item.amount} vendor={item.vendor} arrivalDate={item.arrivalDate} />) : <Box sx={{ display: 'flex' }}> */}
            {productList && productList.map(product => <ItemProduct id={product.sku} sku={product.sku} key={product.sku} images={product.images} name={product.name} itemTypeId={product.item_type_id} desc={product.desc} />)}
            {/* <CircularProgress /> */}
            {/* </Box>} */}
            <Pagination getNumberPerPage={numberPerPageHandle} pageCount={pagination.pageCount} onPageChange={handleClick} />
            {/* <Pagination /> */}
        </div>
    )
}

export default ProductList