import unittest

from ..cell import Board


class TestCellBuilder(unittest.TestCase):

    def test_should_accept_square_board(self):
        Board(['A', 'B', 'C', 'D'])

    def test_should_reject_non_square_boards(self):
        with self.assertRaises(ValueError) as cm:
            Board(['A', 'B', 'C', 'D', 'E'])

        self.assertIn('square', cm.exception.message)

    def test_should_discover_board_size_2x2(self):
        b = Board(['A', 'B', 'C', 'D'])
        self.assertEqual(2, b.board_size())

    def test_should_discover_board_size_3x3(self):
        b = Board(['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I'])
        self.assertEqual(3, b.board_size())

    def test_should_create_cell_matrix_3x3(self):
        b = Board(['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I'])
        self.assertEqual(3, len(b.board()))
        for row_num in range(len(b.board())):
            row = b[row_num]
            self.assertEqual(3, len(row),
                             msg='row {}: should be 3 but was {}'.format(row_num, len(row)))

    def test_should_create_cell_list_3x3(self):
        b = Board(['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I'])
        self.assertEqual(9, len(b.cells()))

    def test_should_resolve_neighbors_2x2(self):
        b = Board(['A', 'B', 'C', 'D'])
        cell = b[0][0]
        self.assertEqual(3, len(cell.neighbors()))

    def test_should_resolve_corner_neighbors_3x3(self):
        b = Board(['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I'])
        cell = b[0][0]
        self.assertEqual(3, len(cell.neighbors()))

    def test_should_resolve_edge_neighbors_3x3(self):
        b = Board(['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I'])
        cell = b[0][1]
        self.assertEqual(5, len(cell.neighbors()))

    def test_should_resolve_center_neighbors_3x3(self):
        b = Board(['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I'])
        cell = b[1][1]
        self.assertEqual(8, len(cell.neighbors()))

    def test_board_should_support_indexing(self):
        b = Board(['A', 'B', 'C', 'D'])
        self.assertEqual('A', b[0][0].value)

