package com.example.BriarfieldPC.demo.entity;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;

import static javax.persistence.GenerationType.IDENTITY;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Builder
@Entity
public class Drive {
    @Id
    @GeneratedValue(strategy = IDENTITY)
    public long Id;
    public String Manufacturer;
    public String Name;
    public Integer Size;
    public Integer Tdp;
    public Integer Price;
    public String technology;
    public String Condition;
}
