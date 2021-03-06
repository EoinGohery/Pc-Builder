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
public class PSU {
    @Id
    @GeneratedValue(strategy = IDENTITY)
    public long Id;
    public String Manufacturer;
    public String Name;
    public Integer Price;
    public Integer Capacity;
    public String Rating;
    public Integer age;

}
