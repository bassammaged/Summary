class Teacher {
    private String designation  = "Teacher";
    private String collegeName  = "Secranix";
    
    public String getDesignation(){
        return designation;
    }

    public String getCollegeName(){
        return collegeName;
    }

    protected void setDesignation(String d){
        this.designation    = d;
    }

    protected void setCollegeName(String c){
        this.collegeName    = c;
    }

    public void does(){
        System.out.println("Teaching.");
    }
}
