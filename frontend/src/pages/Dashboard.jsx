import { useEffect, useState } from "react";

import Navbar from "../components/Navbar";
import UserDialog from "../components/UserDialog";
import UserTable from "../components/UserTable";

import {
  getUsers,
  createUser,
  updateUser,
  deleteUser,
} from "../services/userApi";

function Dashboard() {
  const [users, setUsers] = useState([]);

  const [search, setSearch] = useState("");

  const [dialogOpen, setDialogOpen] = useState(false);

  const [editingUser, setEditingUser] = useState(null);

  const [loading, setLoading] = useState(true);

  async function loadUsers() {
    try {
      setLoading(true);

      const data = await getUsers();

      setUsers(data);
    } catch (error) {
      console.error(error);
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    loadUsers();
  }, []);

  function handleAddUser() {
    setEditingUser(null);
    setDialogOpen(true);
  }

  function handleEdit(user) {
    setEditingUser(user);
    setDialogOpen(true);
  }

  async function handleSubmit(form) {
    try {
      if (editingUser) {
        await updateUser(editingUser._id, form);
      } else {
        await createUser(form);
      }

      setDialogOpen(false);
      setEditingUser(null);

      await loadUsers();
    } catch (error) {
      console.error(error);
      alert("Something went wrong");
    }
  }

  async function handleDelete(id) {
    const confirmDelete = window.confirm(
      "Are you sure you want to delete this user?",
    );

    if (!confirmDelete) {
      return;
    }

    try {
      await deleteUser(id);

      await loadUsers();
    } catch (error) {
      console.error(error);
      alert("Delete failed");
    }
  }

  const filteredUsers = users.filter((user) => {
    const value = search.toLowerCase();

    return (
      user.name?.toLowerCase().includes(value) ||
      user.username?.toLowerCase().includes(value) ||
      user.email?.toLowerCase().includes(value) ||
      user.phoneNo?.includes(value)
    );
  });

  return (
    <>
      <Navbar onAddUser={handleAddUser} />

      <main className="dashboard">
        <div className="dashboard-header">
          <div>
            <h1>Users</h1>

            <p>Manage office employees</p>
          </div>

          <div className="total-card">
            <span>Total Users</span>
            <strong>{users.length}</strong>
          </div>
        </div>

        <div className="search-box">
          <input
            type="text"
            placeholder="Search by name, username, email or phone..."
            value={search}
            onChange={(e) => setSearch(e.target.value)}
          />
        </div>

        {loading ? (
          <div className="loading">Loading users...</div>
        ) : (
          <UserTable
            users={filteredUsers}
            onEdit={handleEdit}
            onDelete={handleDelete}
          />
        )}
      </main>

      <UserDialog
        open={dialogOpen}
        onClose={() => {
          setDialogOpen(false);
          setEditingUser(null);
        }}
        onSubmit={handleSubmit}
        editingUser={editingUser}
      />
    </>
  );
}

export default Dashboard;
