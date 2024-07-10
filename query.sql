select restaurant_id,name,description,price from menu
                                            where deleted_at is null and id=$1  and
                                                  restaurant_id  in ( select id from  restaurants where id=$1)