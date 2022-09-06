import React from 'react';
import Table from 'react-bootstrap/Table'
import axios from 'axios';

class CoffeeList extends React.Component {

    readData() {

        const self = this;
        const instance = axios.create({
            baseURL: 'http://20.115.0.107:9090'
    });
        console.log('20.115.0.107:9090/products')


        instance.get('/products').then(function(response) {
            console.log(response.data);

            self.setState({products: response.data});
        }).catch(function (error){
            console.log(error);
        });
    }

    getProducts() {
        let table = []

        for (let i=0; i < this.state.products.length; i++) {

            table.push(
            <tr key={i}>
                <td>{this.state.products[i].name}</td>
                <td>{this.state.products[i].price}</td>
                <td>{this.state.products[i].sku}</td>
            </tr>
            );
        }

        return table
    }

    constructor(props) {
        super(props);
        this.readData();
        this.state = {products: []};
    
        this.readData = this.readData.bind(this);
    }

    render() {
      return (
        <div>
            <h1 style={{marginBottom: "40px"}}>Menu</h1>
            <Table>
                <thead>
                    <tr>
                        <th>
                            Name
                        </th>
                        <th>
                            Price
                        </th>
                        <th>
                            SKU
                        </th>
                    </tr>
                </thead>
                <tbody>
                    {this.getProducts()}
                </tbody>
            </Table>
        </div>
      ) 
    }
}

export default CoffeeList;