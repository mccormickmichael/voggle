

def mesh_from_board(board):
    node_matrix = [
        [ Node(c) for c in row] for row in board.matrix
    ]
    # build corners
    node_matrix[0][0].add_neighbors([node_matrix[0][1], node_matrix[1][0]])
    node_matrix[0][-1].add_neighbors([node_matrix[1][-1], node_matrix[0][-2]])
    node_matrix[-1][0].add_neighbors([node_matrix[-2][0], node_matrix[-1][1]])
    node_matrix[-1][-1].add_neighbors([node_matrix[-1][-2], node_matrix[-2][-1]])
    
            
    # build edges
    # build body


class Mesh(object):
        
    def __init__(self):
        self.nodes = set()

        


class Node(object):
    def __init__(self, cell):
        self.row_index = cell.row_index
        self.col_index = sell.col_index
        self.value = cell.value
        self.neighbors = set()

    def add_neighbors(self, nodes):
        self.neighbors.update(nodes)

    def __eq__(self, other):
        return self.row_index == other.row_index and self.col_index == other.col_index
    
