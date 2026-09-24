function Navbar({ onAddUser }) {
  return (
    <nav className="navbar">
      <div className="navbar-logo">Office User Management</div>

      <button onClick={onAddUser}>+ Add User</button>
    </nav>
  );
}

export default Navbar;
